package services

import (
	"github.com/pkg/errors"
	"github.com/sbward/authn/data"
)

func PasswordExpirer(store data.AccountStore, tokenStore data.RefreshTokenStore, accountID int) error {
	affected, err := store.RequireNewPassword(accountID)
	if err != nil {
		return errors.Wrap(err, "RequireNewPassword")
	}
	if !affected {
		return FieldErrors{{"account", ErrNotFound}}
	}

	return SessionBatchEnder(tokenStore, accountID)
}
