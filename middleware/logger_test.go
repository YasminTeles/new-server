//nolint:gochecknoglobals,lll
package middleware_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/YasminTeles/new-server/middleware"
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
	request, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/healthcheck", nil)

	requestID := "request-id"
	request.Header.Set("X-Request-ID", requestID)

	logger := middleware.NewLogger()
	logger.Log.Out = &bytes.Buffer{}
	logger.ServeHTTP(recorder, request, func(w http.ResponseWriter, r *http.Request) {})

	lines := strings.Split(strings.TrimSpace(logger.Log.Out.(*bytes.Buffer).String()), "\n")

	expected := fmt.Sprintf(`{"X-Request-ID":"%s","hostname":"","level":"info","method":"GET","msg":"Started handling request","request":"","time":"%s"}`, requestID, nowToday)

	assert.JSONEq(t, expected, lines[0])
}
