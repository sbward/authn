package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	app "github.com/keratin/authn"
	"github.com/keratin/authn/lib/route"
	"github.com/keratin/authn/ops"
	"github.com/keratin/authn/server/cors"
	"github.com/keratin/authn/server/sessions"
)

func Router(app *app.App) http.Handler {
	r := mux.NewRouter()
	route.Attach(r, app.Config.MountedPath, PrivateRoutes(app)...)
	route.Attach(r, app.Config.MountedPath, PublicRoutes(app)...)

	return wrapRouter(r, app)
}

func PublicRouter(app *app.App) http.Handler {
	r := mux.NewRouter()
	route.Attach(r, app.Config.MountedPath, PublicRoutes(app)...)

	return wrapRouter(r, app)
}

func wrapRouter(r *mux.Router, app *app.App) http.Handler {
	stack := handlers.CombinedLoggingHandler(os.Stdout, r)
	stack = sessions.Middleware(app)(stack)
	stack = cors.Middleware(app)(stack)

	if app.Config.Proxied {
		stack = handlers.ProxyHeaders(stack)
	}

	return ops.PanicHandler(app.Reporter, stack)
}
