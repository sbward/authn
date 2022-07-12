package handlers

import (
	"net/http"

	app "github.com/keratin/authn/v2"
	"github.com/keratin/authn/v2/services"
)

func GetAccountsAvailable(app *app.App) http.HandlerFunc {
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
