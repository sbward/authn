package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/sbward/authn/data"
	"github.com/sbward/authn/data/mysql"
	"github.com/sbward/authn/data/postgres"
	dataRedis "github.com/sbward/authn/data/redis"
	"github.com/sbward/authn/data/sqlite3"
	"github.com/sbward/authn/ops"
)

func NewDB(url *url.URL) (*sqlx.DB, error) {
	switch url.Scheme {
	case "sqlite3":
		return sqlite3.NewDB(url.Path)
	case "mysql":
		return mysql.NewDB(url)
	case "postgresql", "postgres":
		return postgres.NewDB(url)
	default:
		return nil, fmt.Errorf("Unsupported database: %s", url.Scheme)
	}
}

func MigrateDB(url *url.URL) error {
	switch url.Scheme {
	case "sqlite3":
		db, err := sqlite3.NewDB(url.Path)
		if err != nil {
			return err
		}
		defer db.Close()

		return sqlite3.MigrateDB(db)
	case "mysql":
		db, err := mysql.NewDB(url)
		if err != nil {
			return err
		}
		defer db.Close()

		return mysql.MigrateDB(db)
	case "postgresql", "postgres":
		db, err := postgres.NewDB(url)
		if err != nil {
			return err
		}
		defer db.Close()

		return postgres.MigrateDB(db)
	default:
		return fmt.Errorf("Unsupported database")
	}
}

func NewAccountStore(db sqlx.Ext) (data.AccountStore, error) {
	switch db.DriverName() {
	case "sqlite3":
		return &sqlite3.AccountStore{Ext: db}, nil
	case "mysql":
		return &mysql.AccountStore{Ext: db}, nil
	case "postgres":
		return &postgres.AccountStore{Ext: db}, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %v", db.DriverName())
	}
}

func NewBlobStore(interval time.Duration, redis *redis.Client, db *sqlx.DB, reporter ops.ErrorReporter) (data.BlobStore, error) {
	// the lifetime of a key should be slightly more than two intervals
	ttl := interval*2 + 10*time.Second

	// the write lock should be greater than the peak time necessary to generate and encrypt a key,
	// plus send it back over the wire to redis. after this time has elapsed, any other authn server
	// may get the lock and generate a competing key.
	lockTime := time.Duration(1500) * time.Millisecond

	if redis != nil {
		return &dataRedis.BlobStore{
			TTL:      ttl,
			LockTime: lockTime,
			Client:   redis,
		}, nil
	}

	switch db.DriverName() {
	case "sqlite3":
		store := &sqlite3.BlobStore{
			TTL:      ttl,
			LockTime: lockTime,
			DB:       db,
		}
		store.Clean(reporter)
		return store, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %v", db.DriverName())
	}
}

func NewRefreshTokenStore(db *sqlx.DB, redis *redis.Client, reporter ops.ErrorReporter, ttl time.Duration) (data.RefreshTokenStore, error) {
	if redis != nil {
		return &dataRedis.RefreshTokenStore{
			Client: redis,
			TTL:    ttl,
		}, nil
	}

	switch db.DriverName() {
	case "sqlite3":
		store := &sqlite3.RefreshTokenStore{
			Ext: db,
			TTL: ttl,
		}
		store.Clean(reporter)
		return store, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %v", db.DriverName())
	}
}
