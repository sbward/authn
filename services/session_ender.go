package services

import (
	"github.com/sbward/authn/data"
	"github.com/sbward/authn/models"
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
