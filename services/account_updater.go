package services

import (
	"strings"

	"github.com/pkg/errors"
	app "github.com/sbward/authn"
	"github.com/sbward/authn/data"
)

func AccountUpdater(store data.AccountStore, cfg *app.Config, accountID int, username string) error {
	username = strings.TrimSpace(username)

	fieldError := UsernameValidator(cfg, username)
	if fieldError != nil {
		return FieldErrors{*fieldError}
	}

	affected, err := store.UpdateUsername(accountID, username)
	if err != nil {
		if data.IsUniquenessError(err) {
			return FieldErrors{{"username", ErrTaken}}
		}

		return errors.Wrap(err, "UpdateUsername")
	}
	if !affected {
		return FieldErrors{{"account", ErrNotFound}}
	}

	return nil
}
