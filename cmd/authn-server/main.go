package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/sbward/authn"
	dataRedis "github.com/sbward/authn/data/redis"
	"github.com/sbward/authn/ops"
	"github.com/sbward/authn/server"
	"github.com/sirupsen/logrus"

	"os"
	"path"
)

// VERSION is a value injected at build time with ldflags
var VERSION string

func main() {
	var cmd string
	if len(os.Args) == 1 {
		cmd = "server"
	} else {
		cmd = os.Args[1]
	}

	cfg, err := authn.ReadEnv()
	if err != nil {
		fmt.Println(err)
		fmt.Println("\nsee: https://github.com/keratin/authn-server/blob/master/docs/config.md")
		return
	}

	if cmd == "server" {
		serve(cfg)
	} else if cmd == "migrate" {
		migrate(cfg)
	} else {
		os.Stderr.WriteString(fmt.Sprintf("unexpected invocation\n"))
		usage()
		os.Exit(2)
	}
}

func serve(cfg *authn.Config) {
	fmt.Println(fmt.Sprintf("~*~ Keratin AuthN v%s ~*~", VERSION))

	// Default logger
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout

	db, err := NewDB(cfg.DatabaseURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	var redis *redis.Client
	if cfg.RedisIsSentinelMode {
		redis, err = dataRedis.NewSentinel(cfg.RedisSentinelMaster, cfg.RedisSentinelNodes, cfg.RedisSentinelPassword)
		if err != nil {
			err = errors.Wrap(err, "redis.NewSentinel")
			fmt.Println(err)
			return
		}
	} else if cfg.RedisURL != nil {
		redis, err = dataRedis.New(cfg.RedisURL)
		if err != nil {
			err = errors.Wrap(err, "redis.New")
			fmt.Println(err)
			return
		}
	}

	errorReporter, err := ops.NewErrorReporter(cfg.ErrorReporterCredentials, cfg.ErrorReporterType, logger)

	accountStore, err := NewAccountStore(db)
	if err != nil {
		err = errors.Wrap(err, "NewAccountStore")
		fmt.Println(err)
		return
	}

	tokenStore, err := NewRefreshTokenStore(db, redis, errorReporter, cfg.RefreshTokenTTL)
	if err != nil {
		err = errors.Wrap(err, "NewRefreshTokenStore")
		fmt.Println(err)
		return
	}

	blobStore, err := NewBlobStore(cfg.AccessTokenTTL, redis, db, errorReporter)
	if err != nil {
		err = errors.Wrap(err, "NewBlobStore")
		fmt.Println(err)
		return
	}

	app, err := authn.NewApp(cfg, db, redis, logger, errorReporter, accountStore, tokenStore, blobStore)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fmt.Sprintf("AUTHN_URL: %s", cfg.AuthNURL))
	fmt.Println(fmt.Sprintf("PORT: %d", cfg.ServerPort))
	if authn.Config.PublicPort != 0 {
		fmt.Println(fmt.Sprintf("PUBLIC_PORT: %d", authn.Config.PublicPort))
	}

	server.Server(app)
}

func migrate(cfg *authn.Config) {
	fmt.Println("Running migrations.")
	err := MigrateDB(cfg.DatabaseURL)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrations complete.")
	}
}

func usage() {
	exe := path.Base(os.Args[0])
	fmt.Println(fmt.Sprintf(`
Usage:
%s server  - run the server (default)
%s migrate - run migrations
`, exe, exe))
}
