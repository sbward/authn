package services

import (
	"github.com/sbward/authn/data"
)

func SessionBatchEnder(store data.RefreshTokenStore, accountID int) error {
	tokens, err := store.FindAll(accountID)
	if err != nil {
		return err
	}
	for _, token := range tokens {
		err = store.Revoke(token)
		if err != nil {
			return err
		}
	}
	return nil
}
