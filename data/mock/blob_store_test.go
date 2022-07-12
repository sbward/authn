package mock_test

import (
	"testing"
	"time"

	"github.com/sbward/authn/data/mock"
	"github.com/sbward/authn/data/testers"
)

func TestBlobStore(t *testing.T) {
	for _, tester := range testers.BlobStoreTesters {
		store := mock.NewBlobStore(time.Second, time.Second)
		tester(t, store)
	}
}
