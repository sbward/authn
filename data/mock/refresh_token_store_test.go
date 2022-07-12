package mock_test

import (
	"testing"

	"github.com/sbward/authn/data/mock"
	"github.com/sbward/authn/data/testers"
)

func TestRefreshTokenStore(t *testing.T) {
	for _, tester := range testers.RefreshTokenStoreTesters {
		store := mock.NewRefreshTokenStore()
		tester(t, store)
	}
}
