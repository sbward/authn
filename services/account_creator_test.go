package services_test

import (
	"testing"

	"github.com/sbward/authn"
	"github.com/sbward/authn/data/mock"
	"github.com/sbward/authn/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountCreatorSuccess(t *testing.T) {
	store := mock.NewAccountStore()

	testCases := []struct {
		config   authn.Config
		username string
		password string
	}{
		{authn.Config{UsernameIsEmail: false, UsernameMinLength: 6}, "userName", "PASSword"},
		{authn.Config{UsernameIsEmail: true}, "username@test.com", "PASSword"},
		{authn.Config{UsernameIsEmail: true, UsernameDomains: []string{"rightdomain.com"}}, "username@rightdomain.com", "PASSword"},
	}

	for _, tc := range testCases {
		acc, err := services.AccountCreator(store, &tc.config, tc.username, tc.password)
		require.NoError(t, err)
		assert.NotEqual(t, 0, acc.ID)
		assert.Equal(t, tc.username, acc.Username)
	}
}

var pw = []byte("$2a$04$ZOBA8E3nT68/ArE6NDnzfezGWEgM6YrE17PrOtSjT5.U/ZGoxyh7e")

func TestAccountCreatorFailure(t *testing.T) {
	store := mock.NewAccountStore()
	store.Create("existing@test.com", pw)

	testCases := []struct {
		config   authn.Config
		username string
		password string
		errors   services.FieldErrors
	}{
		// username validations
		{authn.Config{}, "", "PASSword", services.FieldErrors{{"username", "MISSING"}}},
		{authn.Config{}, "  ", "PASSword", services.FieldErrors{{"username", "MISSING"}}},
		{authn.Config{}, "existing@test.com", "PASSword", services.FieldErrors{{"username", "TAKEN"}}},
		{authn.Config{UsernameIsEmail: true}, "notanemail", "PASSword", services.FieldErrors{{"username", "FORMAT_INVALID"}}},
		{authn.Config{UsernameIsEmail: true}, "@wrong.com", "PASSword", services.FieldErrors{{"username", "FORMAT_INVALID"}}},
		{authn.Config{UsernameIsEmail: true}, "wrong@wrong", "PASSword", services.FieldErrors{{"username", "FORMAT_INVALID"}}},
		{authn.Config{UsernameIsEmail: true}, "wrong@wrong.", "PASSword", services.FieldErrors{{"username", "FORMAT_INVALID"}}},
		{authn.Config{UsernameIsEmail: true, UsernameDomains: []string{"rightdomain.com"}}, "email@wrongdomain.com", "PASSword", services.FieldErrors{{"username", "FORMAT_INVALID"}}},
		{authn.Config{UsernameIsEmail: false, UsernameMinLength: 6}, "short", "PASSword", services.FieldErrors{{"username", "FORMAT_INVALID"}}},
		// password validations
		{authn.Config{}, "username", "", services.FieldErrors{{"password", "MISSING"}}},
		{authn.Config{PasswordMinComplexity: 2}, "username", "qwerty", services.FieldErrors{{"password", "INSECURE"}}},
		{authn.Config{UsernameIsEmail: true}, "username@test.example.com", "username@test.example.com", services.FieldErrors{{"password", "INSECURE"}}},
	}

	for _, tc := range testCases {
		t.Run(tc.username, func(t *testing.T) {
			acc, err := services.AccountCreator(store, &tc.config, tc.username, tc.password)
			if assert.Equal(t, tc.errors, err) {
				assert.Empty(t, acc)
			}
		})
	}
}
