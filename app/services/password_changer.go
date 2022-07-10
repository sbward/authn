package services

import (
	app "github.com/keratin/authn"
	"github.com/keratin/authn/data"
	"github.com/keratin/authn/ops"
	"github.com/pkg/errors"
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
