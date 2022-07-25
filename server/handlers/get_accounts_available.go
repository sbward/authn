package handlers

import (
	"net/http"

	"github.com/sbward/authn"
	"github.com/sbward/authn/services"
)

func GetAccountsAvailable(app *authn.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		account, err := app.AccountStore.FindByUsername(r.FormValue("username"))
		if err != nil {
			panic(err)
		}

		if account == nil {
			WriteData(w, http.StatusOK, true)
		} else {
			WriteErrors(w, services.FieldErrors{{"username", services.ErrTaken}})
		}
	}
}
