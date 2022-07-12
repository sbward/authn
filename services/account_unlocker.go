package services

import (
	"github.com/pkg/errors"
	"github.com/sbward/authn/data"
)

func AccountUnlocker(store data.AccountStore, accountID int) error {
	affected, err := store.Unlock(accountID)
	if err != nil {
		return errors.Wrap(err, "Unlock")
	}
	if !affected {
		return FieldErrors{{"account", ErrNotFound}}
	}

	return nil
}
