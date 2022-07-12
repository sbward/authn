package data_test

import (
	"testing"
	"time"

	"github.com/keratin/authn/v2/data"
	"github.com/keratin/authn/v2/data/mock"
	"github.com/stretchr/testify/assert"
)

func TestEncryptedBlobStore(t *testing.T) {
	bs := mock.NewBlobStore(time.Second, time.Second)
	ebs := data.NewEncryptedBlobStore(bs, []byte("secretsecretsecretsecretsecret12"))
	val := []byte("val")

	ok, err := ebs.WriteNX("key", val)
	assert.NoError(t, err)
	assert.True(t, ok)

	blob, err := bs.Read("key")
	assert.NoError(t, err)
	assert.NotEmpty(t, blob)
	assert.NotEqual(t, val, blob)

	blob, err = ebs.Read("key")
	assert.NoError(t, err)
	assert.Equal(t, val, blob)
}
