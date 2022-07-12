package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/keratin/authn/v2/data/redis"
	"github.com/keratin/authn/v2/data/testers"
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
