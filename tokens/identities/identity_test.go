package identities_test

import (
	"net/url"
	"testing"

	app "github.com/sbward/authn"
	"github.com/sbward/authn/data/private"

	"gopkg.in/square/go-jose.v2"

	"github.com/sbward/authn/data/mock"
	"github.com/sbward/authn/tokens/identities"
	"github.com/sbward/authn/tokens/sessions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIdentityClaims(t *testing.T) {
	store := mock.NewRefreshTokenStore()
	cfg := app.Config{
		AuthNURL:          &url.URL{Scheme: "http", Host: "authn.example.com"},
		SessionSigningKey: []byte("key-a-reno"),
	}
	key, err := private.GenerateKey(512)
	require.NoError(t, err)
	session, err := sessions.New(store, &cfg, 1, "example.com")
	require.NoError(t, err)

	t.Run("includes KID", func(t *testing.T) {
		identity := identities.New(&cfg, session, 1, "example.com")
		identityStr, err := identity.Sign(key)
		require.NoError(t, err)

		parsed, err := jose.ParseSigned(identityStr)
		require.NoError(t, err)

		require.NoError(t, err)
		assert.Equal(t, key.JWK.KeyID, parsed.Signatures[0].Header.KeyID)
	})
}
