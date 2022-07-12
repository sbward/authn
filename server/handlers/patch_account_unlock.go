package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	app "github.com/keratin/authn/v2"
	"github.com/keratin/authn/v2/services"
)

func PatchAccountUnlock(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			WriteNotFound(w, "account")
			return
		}

		err = services.AccountUnlocker(app.AccountStore, id)
		if err != nil {
			if _, ok := err.(services.FieldErrors); ok {
				WriteNotFound(w, "account")
				return
			}

			panic(err)
		}

		w.WriteHeader(http.StatusOK)
	}
}
