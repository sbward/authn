package data

import (
	"github.com/sbward/authn/models"
)

type AccountStore interface {
	Create(u string, p []byte) (*models.Account, error)
	Find(id int) (*models.Account, error)
	FindByUsername(u string) (*models.Account, error)
	FindByOauthAccount(p string, pid string) (*models.Account, error)
	AddOauthAccount(id int, p string, pid string, tok string) error
	GetOauthAccounts(id int) ([]*models.OauthAccount, error)
	Archive(id int) (bool, error)
	Lock(id int) (bool, error)
	Unlock(id int) (bool, error)
	RequireNewPassword(id int) (bool, error)
	SetPassword(id int, p []byte) (bool, error)
	UpdateUsername(id int, u string) (bool, error)
	SetLastLogin(id int) (bool, error)
}
