package data

import (
	"github.com/keratin/authn/v2/data/private"
)

type KeyStore interface {
	// Returns the current key
	Key() *private.Key
	// Returns recent keys (including current key)
	Keys() []*private.Key
}
