package services

import (
	"github.com/pkg/errors"
	"github.com/sbward/authn/data"
)

func AccountLocker(store data.AccountStore, tokenStore data.RefreshTokenStore, accountID int) error {
	affected, err := store.Lock(accountID)
	if err != nil {
		return errors.Wrap(err, "Lock")
	}
	if !affected {
		return FieldErrors{{"account", ErrNotFound}}
	}

	return SessionBatchEnder(tokenStore, accountID)
}
