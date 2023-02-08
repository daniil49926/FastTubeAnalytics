package api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_handleHealthChecker(t *testing.T) {
	server := New(NewConfig())
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/healthchecker", nil)
	server.handleHealthChecker().ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Body.String(), "{\"result\":\"OK\"}")
}
