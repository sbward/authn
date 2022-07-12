package services_test

import (
	"testing"

	"github.com/keratin/authn/v2/data/mock"
	"github.com/keratin/authn/v2/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSessionBatchEnder(t *testing.T) {
	store := mock.NewRefreshTokenStore()

	t.Run("revoking nothing", func(t *testing.T) {
		id := 123
		err := services.SessionBatchEnder(store, id)
		assert.NoError(t, err)
	})

	t.Run("revoking something", func(t *testing.T) {
		id := 234
		_, err := store.Create(id)
		require.NoError(t, err)

		found, err := store.FindAll(id)
		require.NoError(t, err)
		require.Len(t, found, 1)

		err = services.SessionBatchEnder(store, id)
		assert.NoError(t, err)

		found, err = store.FindAll(id)
		assert.NoError(t, err)
		assert.Len(t, found, 0)
	})
}
