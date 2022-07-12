package services

import (
	"strings"

	"github.com/pkg/errors"
	app "github.com/sbward/authn"
	"github.com/sbward/authn/data"
	"github.com/sbward/authn/models"
	"golang.org/x/crypto/bcrypt"
)

func AccountCreator(store data.AccountStore, cfg *app.Config, username string, password string) (*models.Account, error) {
	username = strings.TrimSpace(username)

	errs := FieldErrors{}

	fieldError := UsernameValidator(cfg, username)
	if fieldError != nil {
		errs = append(errs, *fieldError)
	}

	fieldError = PasswordValidator(cfg, username, password)
	if fieldError != nil {
		errs = append(errs, *fieldError)
	}

	if len(errs) > 0 {
		return nil, errs
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), cfg.BcryptCost)
	if err != nil {
		return nil, errors.Wrap(err, "bcrypt")
	}

	acc, err := store.Create(username, hash)
	if err != nil {
		if data.IsUniquenessError(err) {
			return nil, FieldErrors{{"username", ErrTaken}}
		}

		return nil, errors.Wrap(err, "Create")
	}

	return acc, nil
}
