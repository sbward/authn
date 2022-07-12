package handlers_test

import (
	"net/http"
	"testing"

	"github.com/keratin/authn/v2/lib/route"
	"github.com/keratin/authn/v2/server/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRoot(t *testing.T) {
	app := test.App()
	server := test.Server(app)
	defer server.Close()

	client := route.NewClient(server.URL)

	res, err := client.Get("/")
	require.NoError(t, err)
	body := test.ReadBody(res)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, []string{"text/html"}, res.Header["Content-Type"])
	assert.NotEmpty(t, body)
}
