package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"
)

var (
	nowTime  = time.Now()
	nowToday = nowTime.Format("02-01-2006 15:04:05")
)

func TestLogger(t *testing.T) {
	t.Parallel()

	recorder := negroni.NewResponseWriter(httptest.NewRecorder())
	request, _ := http.NewRequest("GET", "/healthcheck", nil)

	requestID := "request-id"
	request.Header.Set("X-Request-ID", requestID)

	logger := NewLogger()
	logger.log.Out = &bytes.Buffer{}
	logger.ServeHTTP(recorder, request, func(w http.ResponseWriter, r *http.Request) {})

	lines := strings.Split(strings.TrimSpace(logger.log.Out.(*bytes.Buffer).String()), "\n")

	expected := fmt.Sprintf(`{"X-Request-ID":"%s","hostname":"","level":"info","method":"GET","msg":"Started handling request","request":"","time":"%s"}`, requestID, nowToday)

	assert.JSONEq(t, expected, lines[0])
}
