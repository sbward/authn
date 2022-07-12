package mock_test

import (
	"testing"

	"github.com/keratin/authn/v2/data/mock"
	"github.com/keratin/authn/v2/data/testers"
)

func TestRefreshTokenStore(t *testing.T) {
	for _, tester := range testers.RefreshTokenStoreTesters {
		store := mock.NewRefreshTokenStore()
		tester(t, store)
	}
}
