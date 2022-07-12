package handlers

import (
	"bytes"
	"net/http"

	app "github.com/sbward/authn"
	"github.com/sbward/authn/server/views"
)

func GetRoot(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		views.Root(&buf)

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write(buf.Bytes())
	}
}
