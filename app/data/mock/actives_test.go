package mock_test

import (
	"testing"

	"github.com/keratin/authn/data/mock"
	"github.com/keratin/authn/data/testers"
)

func TestActives(t *testing.T) {
	for _, tester := range testers.ActivesTesters {
		mStore := mock.NewActives()
		tester(t, mStore)
	}
}
