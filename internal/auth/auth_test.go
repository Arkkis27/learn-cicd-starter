package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "ApiKey my-secret-api-key")
	apiKey, err := GetAPIKey(headers)
	assert.NoError(t, err)
	assert.Equal(t, "my-secret-api-key", apiKey)
}

func TestAuthMissing(t *testing.T) {
	headers := make(http.Header)
	apiKey, err := GetAPIKey(headers)
	assert.Error(t, err)
	assert.Equal(t, "", apiKey)
}

func TestAuthErr(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "ApiKey")
	apiKey, err := GetAPIKey(headers)
	assert.Error(t, err)
	assert.Equal(t, "", apiKey)
}

func TestAuthFormat(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "ApiKey: my-secret-api-key")
	apiKey, err := GetAPIKey(headers)
	assert.Error(t, err)
	assert.Equal(t, "", apiKey)
}

func TestAuthKeyFormat(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "ApiKey my-secret api-key")
	apiKey, err := GetAPIKey(headers)
	assert.Error(t, err)
	assert.Equal(t, "my-secret", apiKey)
}
