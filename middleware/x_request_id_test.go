package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const XRequestIDHeader = "X-Request-ID"

func TestNonExistXRequestIDInHeader(t *testing.T) {
	t.Parallel()

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)

	xRequestID := NewXRequestID()
	xRequestID.ServeHTTP(recorder, req, func(w http.ResponseWriter, r *http.Request) {})

	assert.NotEmpty(t, req.Header.Get(XRequestIDHeader))
	assert.NotEmpty(t, recorder.HeaderMap.Get(XRequestIDHeader))

	_, err := uuid.Parse(req.Header.Get(XRequestIDHeader))
	assert.NoError(t, err)

	_, err = uuid.Parse(recorder.HeaderMap.Get(XRequestIDHeader))
	assert.NoError(t, err)
}

func TestExistXRequestIDInHeader(t *testing.T) {
	t.Parallel()

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)

	xRequestIDValue := "test-id"
	req.Header.Set(XRequestIDHeader, xRequestIDValue)

	xRequestID := NewXRequestID()
	xRequestID.ServeHTTP(recorder, req, func(w http.ResponseWriter, r *http.Request) {})

	assert.Equal(t, xRequestIDValue, req.Header.Get(XRequestIDHeader))
	assert.Equal(t, xRequestIDValue, recorder.HeaderMap.Get(XRequestIDHeader))
}
