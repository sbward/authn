package test

import (
	"net/http"
	"net/url"

	app "github.com/keratin/authn/v2"
	"github.com/keratin/authn/v2/data/mock"
	"github.com/keratin/authn/v2/data/private"
	"github.com/keratin/authn/v2/lib/oauth"
	"github.com/keratin/authn/v2/lib/route"
	"github.com/keratin/authn/v2/ops"
	"github.com/sirupsen/logrus"
)

func App() *app.App {
	authnURL, err := url.Parse("https://authn.example.com")
	if err != nil {
		panic(err)
	}

	weakKey, err := private.GenerateKey(512)
	if err != nil {
		panic(err)
	}

	cfg := app.Config{
		BcryptCost:              4,
		SessionSigningKey:       []byte("TestKey"),
		AuthNURL:                authnURL,
		SessionCookieName:       "authn",
		OAuthCookieName:         "authn-oauth-nonce",
		ApplicationDomains:      []route.Domain{{Hostname: "test.com"}},
		PasswordMinComplexity:   2,
		AppPasswordResetURL:     &url.URL{Scheme: "https", Host: "app.example.com"},
		AppPasswordlessTokenURL: &url.URL{Scheme: "https", Host: "app.example.com"},
		EnableSignup:            true,
		SameSite:                http.SameSiteDefaultMode,
		PasswordChangeLogout:    false,
	}

	logger := logrus.New()
	return &app.App{
		Config:            &cfg,
		KeyStore:          mock.NewKeyStore(weakKey),
		AccountStore:      mock.NewAccountStore(),
		RefreshTokenStore: mock.NewRefreshTokenStore(),
		Actives:           mock.NewActives(),
		Reporter:          &ops.LogReporter{logger},
		OauthProviders:    map[string]oauth.Provider{},
		Logger:            logger,
	}
}
