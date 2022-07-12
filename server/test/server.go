package test

import (
	"net/http/httptest"

	app "github.com/keratin/authn/v2"
	"github.com/keratin/authn/v2/server"
)

func Server(app *app.App) *httptest.Server {
	return httptest.NewServer(server.Router(app))
}
