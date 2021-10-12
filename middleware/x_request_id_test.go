package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

  "github.com/google/uuid"
)

func TestNonExistXRequestIDInHeader(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)

  xRequestID := NewXRequestID()
  xRequestID.ServeHTTP(recorder, req, func(w http.ResponseWriter, r *http.Request) {})

	if requestID := req.Header.Get("X-Request-ID"); requestID == "" {
		t.Fatalf("Expected request X-Request-Id not to be empty, got `%v`", requestID)
	}

	if responseID := recorder.HeaderMap.Get("X-Request-ID"); responseID == "" {
		t.Fatalf("Expected response X-Request-Id not to be empty, got `%v`", responseID)
	}

  if _,err := uuid.Parse(req.Header.Get("X-Request-ID")); err != nil {
    t.Fatalf("Expected request X-Request-Id to be an uuid")
  }

  if _,err := uuid.Parse(recorder.HeaderMap.Get("X-Request-ID")); err != nil {
    t.Fatalf("Expected response X-Request-Id to be an uuid")
  }
}

func TestExistXRequestIDInHeader(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
  req.Header.Set("X-Request-ID", "test-id")

  xRequestID := NewXRequestID()
  xRequestID.ServeHTTP(recorder, req, func(w http.ResponseWriter, r *http.Request) {})

	if requestID := req.Header.Get("X-Request-ID"); requestID != "test-id" {
		t.Fatalf("Expected request X-Request-Id to be `test-id`, got `%v`", requestID)
	}

	if responseID := recorder.HeaderMap.Get("X-Request-ID"); responseID != "test-id" {
		t.Fatalf("Expected response X-Request-Id to be `test-id`, got `%v`", responseID)
	}
}
