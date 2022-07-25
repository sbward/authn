package handlers_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/sbward/authn"
	"github.com/sbward/authn/data/private"
	"github.com/sirupsen/logrus"

	"github.com/sbward/authn/data/mock"
	"github.com/sbward/authn/server/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetJWKs(t *testing.T) {
	rsaKey, err := private.GenerateKey(512)
	require.NoError(t, err)
	app := &authn.App{
		KeyStore: mock.NewKeyStore(rsaKey),
		Config:   &authn.Config{},
		Logger:   logrus.New(),
	}

	server := test.Server(app)
	defer server.Close()

	res, err := http.Get(fmt.Sprintf("%s/jwks", server.URL))
	require.NoError(t, err)
	body := test.ReadBody(res)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, []string{"application/json"}, res.Header["Content-Type"])
	assert.NotEmpty(t, body)
}

func BenchmarkGetJWKs(b *testing.B) {
	rsaKey, _ := private.GenerateKey(2048)
	app := &authn.App{
		KeyStore: mock.NewKeyStore(rsaKey),
		Config:   &authn.Config{},
	}

	server := test.Server(app)
	defer server.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		http.Get(fmt.Sprintf("%s/jwks", server.URL))
	}
}
