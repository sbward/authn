package data

import "errors"

var ErrNotUnique = errors.New("Value is not unique")

func IsUniquenessError(err error) bool {
	return err == ErrNotUnique
}
