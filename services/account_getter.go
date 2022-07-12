package services

import (
	"github.com/pkg/errors"
	"github.com/sbward/authn/data"
	"github.com/sbward/authn/models"
)

func AccountGetter(store data.AccountStore, accountID int) (*models.Account, error) {
	account, err := store.Find(accountID)
	if err != nil {
		return nil, errors.Wrap(err, "Find")
	}
	if account == nil {
		return nil, FieldErrors{{"account", ErrNotFound}}
	}

	return account, nil
}
