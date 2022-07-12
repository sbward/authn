package handlers

import (
	"net/http"
	"regexp"

	app "github.com/sbward/authn"
	"github.com/sbward/authn/lib/parse"

	"github.com/sbward/authn/services"
)

func PostAccountsImport(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user struct {
			Username string
			Password string
			Locked   string
		}
		if err := parse.Payload(r, &user); err != nil {
			WriteErrors(w, err)
			return
		}
		locked, err := regexp.MatchString("^(?i:t|true|yes)$", user.Locked)
		if err != nil {
			panic(err)
		}

		account, err := services.AccountImporter(
			app.AccountStore,
			app.Config,
			user.Username,
			user.Password,
			locked,
		)
		if err != nil {
			if fe, ok := err.(services.FieldErrors); ok {
				WriteErrors(w, fe)
				return
			}

			panic(err)
		}

		WriteData(w, http.StatusCreated, map[string]int{
			"id": account.ID,
		})
	}
}
