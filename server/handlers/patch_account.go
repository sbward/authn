package handlers

import (
	"net/http"
	"strconv"

	app "github.com/keratin/authn"
	"github.com/keratin/authn/lib/parse"

	"github.com/gorilla/mux"
	"github.com/keratin/authn/services"
)

func PatchAccount(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user struct{ Username string }
		if err := parse.Payload(r, &user); err != nil {
			WriteErrors(w, err)
			return
		}
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			WriteNotFound(w, "account")
			return
		}

		err = services.AccountUpdater(app.AccountStore, app.Config, id, user.Username)
		if err != nil {
			if fe, ok := err.(services.FieldErrors); ok {
				if fe[0].Message == services.ErrNotFound {
					WriteNotFound(w, "account")
				} else {
					WriteErrors(w, fe)
				}
				return
			}

			panic(err)
		}

		w.WriteHeader(http.StatusOK)
	}
}
