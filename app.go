package authn

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sbward/authn/data"
	"github.com/sbward/authn/lib/oauth"
	"github.com/sbward/authn/ops"
	"github.com/sirupsen/logrus"

	dataRedis "github.com/sbward/authn/data/redis"
)

type pinger func() bool

type App struct {
	DB                *sqlx.DB
	DbCheck           pinger
	RedisCheck        pinger
	Config            *Config
	AccountStore      data.AccountStore
	RefreshTokenStore data.RefreshTokenStore
	KeyStore          data.KeyStore
	Actives           data.Actives
	Reporter          ops.ErrorReporter
	OauthProviders    map[string]oauth.Provider
	Logger            logrus.FieldLogger
}

func NewApp(cfg *Config, db *sqlx.DB, redis *redis.Client, logger logrus.FieldLogger, errorReporter ops.ErrorReporter, accountStore data.AccountStore, tokenStore data.RefreshTokenStore, blobStore data.BlobStore) (*App, error) {

	keyStore := data.NewRotatingKeyStore()
	if cfg.IdentitySigningKey == nil {
		m := data.NewKeyStoreRotater(
			data.NewEncryptedBlobStore(blobStore, cfg.DBEncryptionKey),
			cfg.AccessTokenTTL,
			logger,
		)
		err := m.Maintain(keyStore, errorReporter)
		if err != nil {
			return nil, errors.Wrap(err, "Maintain")
		}
	} else {
		keyStore.Rotate(cfg.IdentitySigningKey)
	}

	var actives data.Actives
	if redis != nil {
		actives = dataRedis.NewActives(
			redis,
			cfg.StatisticsTimeZone,
			cfg.DailyActivesRetention,
			cfg.WeeklyActivesRetention,
			5*12,
		)
	}

	oauthProviders := map[string]oauth.Provider{}
	if cfg.GoogleOauthCredentials != nil {
		oauthProviders["google"] = *oauth.NewGoogleProvider(cfg.GoogleOauthCredentials)
	}
	if cfg.GitHubOauthCredentials != nil {
		oauthProviders["github"] = *oauth.NewGitHubProvider(cfg.GitHubOauthCredentials)
	}
	if cfg.FacebookOauthCredentials != nil {
		oauthProviders["facebook"] = *oauth.NewFacebookProvider(cfg.FacebookOauthCredentials)
	}
	if cfg.DiscordOauthCredentials != nil {
		oauthProviders["discord"] = *oauth.NewDiscordProvider(cfg.DiscordOauthCredentials)
	}
	if cfg.MicrosoftOauthCredientials != nil {
		oauthProviders["microsoft"] = *oauth.NewMicrosoftProvider(cfg.MicrosoftOauthCredientials)
	}

	return &App{
		// Provide access to root DB - useful when extending AccountStore functionality
		DB:                db,
		DbCheck:           func() bool { return db.Ping() == nil },
		RedisCheck:        func() bool { return redis != nil && redis.Ping(context.TODO()).Err() == nil },
		Config:            cfg,
		AccountStore:      accountStore,
		RefreshTokenStore: tokenStore,
		KeyStore:          keyStore,
		Actives:           actives,
		Reporter:          errorReporter,
		OauthProviders:    oauthProviders,
		Logger:            logger,
	}, nil
}
