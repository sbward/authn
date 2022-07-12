package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/sbward/authn/data/redis"
	"github.com/sbward/authn/data/testers"
	"github.com/stretchr/testify/require"
)

func TestRefreshTokenStore(t *testing.T) {
	client, err := redis.TestDB()
	require.NoError(t, err)
	store := &redis.RefreshTokenStore{Client: client, TTL: time.Second}
	for _, tester := range testers.RefreshTokenStoreTesters {
		tester(t, store)
		store.FlushDB(context.TODO())
	}
}
