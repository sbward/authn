package services

import (
	"github.com/pkg/errors"
	app "github.com/sbward/authn"
	"github.com/sbward/authn/data"
	"github.com/sbward/authn/ops"
	"golang.org/x/crypto/bcrypt"
)

func PasswordChanger(store data.AccountStore, r ops.ErrorReporter, cfg *app.Config, id int, currentPassword string, password string) error {
	account, err := store.Find(id)
	if err != nil {
		return errors.Wrap(err, "Find")
	}
	if account == nil {
		return FieldErrors{{"account", ErrNotFound}}
	} else if account.Locked {
		return FieldErrors{{"account", ErrLocked}}
	}

	err = bcrypt.CompareHashAndPassword(account.Password, []byte(currentPassword))
	if err != nil {
		return FieldErrors{{"credentials", ErrFailed}}
	}

	return PasswordSetter(store, r, cfg, id, password)
}
