package middleware_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/YasminTeles/new-server/middleware"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const XRequestIDHeader = "X-Request-ID"

func TestNonExistXRequestIDInHeader(t *testing.T) {
	t.Parallel()

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/healthcheck", nil)

	xRequestID := middleware.NewXRequestID()
	xRequestID.ServeHTTP(recorder, request, func(w http.ResponseWriter, r *http.Request) {})

	assert.NotEmpty(t, request.Header.Get(XRequestIDHeader))
	assert.NotEmpty(t, recorder.Header().Get(XRequestIDHeader))

	_, err := uuid.Parse(request.Header.Get(XRequestIDHeader))
	assert.NoError(t, err)

	_, err = uuid.Parse(recorder.Header().Get(XRequestIDHeader))
	assert.NoError(t, err)
}

func TestExistXRequestIDInHeader(t *testing.T) {
	t.Parallel()

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/healthcheck", nil)

	xRequestIDValue := "test-id"
	req.Header.Set(XRequestIDHeader, xRequestIDValue)

	xRequestID := middleware.NewXRequestID()
	xRequestID.ServeHTTP(recorder, req, func(w http.ResponseWriter, r *http.Request) {})

	assert.Equal(t, xRequestIDValue, req.Header.Get(XRequestIDHeader))
	assert.Equal(t, xRequestIDValue, recorder.Header().Get(XRequestIDHeader))
}
