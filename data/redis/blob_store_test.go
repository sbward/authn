package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/sbward/authn/data/redis"
	"github.com/sbward/authn/data/testers"
	"github.com/stretchr/testify/require"
)

func TestBlobStore(t *testing.T) {
	client, err := redis.TestDB()
	require.NoError(t, err)
	store := &redis.BlobStore{
		Client:   client,
		TTL:      time.Second,
		LockTime: time.Second,
	}
	for _, tester := range testers.BlobStoreTesters {
		tester(t, store)
		client.FlushDB(context.TODO())
	}
}
