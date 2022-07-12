package services

import (
	"strconv"

	app "github.com/sbward/authn"
	"github.com/sbward/authn/ops"

	"github.com/pkg/errors"
	"github.com/sbward/authn/data"
	"github.com/sbward/authn/tokens/resets"
)

func PasswordResetter(store data.AccountStore, r ops.ErrorReporter, cfg *app.Config, token string, password string) (int, error) {
	claims, err := resets.Parse(token, cfg)
	if err != nil {
		return 0, FieldErrors{{"token", ErrInvalidOrExpired}}
	}

	id, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, errors.Wrap(err, "Atoi")
	}

	account, err := store.Find(id)
	if err != nil {
		return 0, errors.Wrap(err, "Find")
	}
	if account == nil {
		return 0, FieldErrors{{"account", ErrNotFound}}
	} else if account.Locked {
		return 0, FieldErrors{{"account", ErrLocked}}
	} else if account.Archived() {
		return 0, FieldErrors{{"account", ErrLocked}}
	}

	if claims.LockExpired(account.PasswordChangedAt) {
		return 0, FieldErrors{{"token", ErrInvalidOrExpired}}
	}

	return account.ID, PasswordSetter(store, r, cfg, id, password)
}
