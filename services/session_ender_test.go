package services_test

import (
	"testing"

	"github.com/sbward/authn/data/mock"
	"github.com/sbward/authn/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSessionEnder(t *testing.T) {
	accountID := 0
	refreshStore := mock.NewRefreshTokenStore()

	t.Run("revokes token", func(t *testing.T) {
		token, err := refreshStore.Create(accountID)
		require.NoError(t, err)

		err = services.SessionEnder(refreshStore, &token)
		assert.NoError(t, err)

		foundID, err := refreshStore.Find(token)
		assert.Empty(t, foundID)
		assert.NoError(t, err)
	})

	t.Run("ignores missing token", func(t *testing.T) {
		err := services.SessionEnder(refreshStore, nil)
		assert.NoError(t, err)
	})
}
