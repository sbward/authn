package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/keratin/authn/v2/data/redis"
	"github.com/keratin/authn/v2/data/testers"
	"github.com/stretchr/testify/require"
)

func TestActives(t *testing.T) {
	client, err := redis.TestDB()
	require.NoError(t, err)
	rStore := redis.NewActives(client, time.UTC, 365, 52, 12)
	for _, tester := range testers.ActivesTesters {
		client.FlushDB(context.TODO())
		tester(t, rStore)
	}
}
