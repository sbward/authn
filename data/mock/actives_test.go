package mock_test

import (
	"testing"

	"github.com/sbward/authn/data/mock"
	"github.com/sbward/authn/data/testers"
)

func TestActives(t *testing.T) {
	for _, tester := range testers.ActivesTesters {
		mStore := mock.NewActives()
		tester(t, mStore)
	}
}
