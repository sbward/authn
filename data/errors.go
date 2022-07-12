package data

import (
	"errors"
	"fmt"
)

type UniquenessError struct {
	Wrap error
}

var _ error = (*UniquenessError)(nil)

func (u UniquenessError) Error() string {
	return fmt.Sprintf("Uniqueness error: %s", u.Wrap.Error())
}

func (u UniquenessError) Unwrap() error {
	return u.Wrap
}

func NewUniquenessError(err error) error {
	return UniquenessError{
		Wrap: err,
	}
}

func IsUniquenessError(err error) bool {
	notUnique := &UniquenessError{}
	return errors.As(err, notUnique)
}
