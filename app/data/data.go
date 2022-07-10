package data

import (
	"fmt"
	"net/url"

	"github.com/lib/pq"

	my "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/keratin/authn/data/mock"
	"github.com/keratin/authn/data/mysql"
	"github.com/keratin/authn/data/postgres"
	"github.com/keratin/authn/data/sqlite3"
	sq3 "github.com/mattn/go-sqlite3"
)

type Migrator interface {
	MigrateDB(db *sqlx.DB) error
}

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

func IsUniquenessError(err error) bool {
	switch i := err.(type) {
	case sq3.Error:
		return i.ExtendedCode == sq3.ErrConstraintUnique
	case *my.MySQLError:
		return i.Number == 1062
	case *pq.Error:
		return i.Code.Class().Name() == "integrity_constraint_violation"
	case mock.Error:
		return i.Code == mock.ErrNotUnique
	default:
		return false
	}
}
