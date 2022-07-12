package services

import (
	"github.com/keratin/authn/v2/data"
	"github.com/keratin/authn/v2/models"
)

func SessionEnder(
	refreshTokenStore data.RefreshTokenStore,
	existingToken *models.RefreshToken,
) (err error) {
	if existingToken != nil {
		return refreshTokenStore.Revoke(*existingToken)
	}
	return nil
}
