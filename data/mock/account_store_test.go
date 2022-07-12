package mock_test

import (
	"testing"

	"github.com/sbward/authn/data/mock"
	"github.com/sbward/authn/data/testers"
)

func TestAccountStore(t *testing.T) {
	for _, tester := range testers.AccountStoreTesters {
		store := mock.NewAccountStore()
		tester(t, store)
	}
}
