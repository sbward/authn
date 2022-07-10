package handlers

import (
	"net/http"

	app "github.com/keratin/authn"
	"github.com/keratin/authn/services"
)

func GetSessionToken(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		account, err := app.AccountStore.FindByUsername(r.FormValue("username"))
		if err != nil {
			panic(err)
		}

		// run in the background so that a timing attack can't enumerate usernames
		go func() {
			err := services.PasswordlessTokenSender(app.Config, account, app.Logger)
			if err != nil {
				app.Reporter.ReportRequestError(err, r)
			}
		}()

		w.WriteHeader(http.StatusOK)
	}
}
