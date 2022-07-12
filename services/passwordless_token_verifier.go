package services

import (
	"strconv"

	app "github.com/sbward/authn"
	"github.com/sbward/authn/ops"

	"github.com/pkg/errors"
	"github.com/sbward/authn/data"
	"github.com/sbward/authn/tokens/passwordless"
)

func PasswordlessTokenVerifier(store data.AccountStore, r ops.ErrorReporter, cfg *app.Config, token string) (int, error) {
	claims, err := passwordless.Parse(token, cfg)
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
	} else if account.LastLoginAt != nil && account.LastLoginAt.After(claims.IssuedAt.Time()) {
		return 0, FieldErrors{{"token", ErrInvalidOrExpired}}
	}

	return account.ID, nil
}
