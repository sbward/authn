package services_test

import (
	"net/url"
	"testing"

	app "github.com/sbward/authn"
	"github.com/sbward/authn/data/mock"
	"github.com/sbward/authn/data/private"
	"github.com/sbward/authn/lib/route"
	"github.com/sbward/authn/ops"
	"github.com/sbward/authn/services"
	"github.com/sbward/authn/tokens/sessions"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSessionRefresher(t *testing.T) {
	rsaKey, err := private.GenerateKey(512)
	require.NoError(t, err)
	keyStore := mock.NewKeyStore(rsaKey)
	cfg := &app.Config{
		AuthNURL: &url.URL{Scheme: "http", Host: "authn.example.com"},
	}
	refreshStore := mock.NewRefreshTokenStore()
	reporter := &ops.LogReporter{logrus.New()}

	accountID := 0
	audience := &route.Domain{"authn.example.com", "8080"}
	session, err := sessions.New(refreshStore, cfg, accountID, audience.String())
	require.NoError(t, err)

	t.Run("tracks actives while generating token", func(t *testing.T) {
		activesStore := mock.NewActives()

		identityToken, err := services.SessionRefresher(
			refreshStore, keyStore, activesStore, cfg, reporter,
			session, accountID, audience,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, identityToken)

		report, err := activesStore.ActivesByDay()
		require.NoError(t, err)
		assert.Len(t, report, 1)
	})

	t.Run("ignores actives when not configured", func(t *testing.T) {
		identityToken, err := services.SessionRefresher(
			refreshStore, keyStore, nil, cfg, reporter,
			session, accountID, audience,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, identityToken)
	})
}
