//nolint:exhaustivestruct
package server_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheckEndPoint(t *testing.T) {
	t.Parallel()

	port := 3000
	url := fmt.Sprintf("http://localhost:%d/healthcheck", port)
	request, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)

	client := http.Client{}
	response, _ := client.Do(request)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	byteBody, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	message := strings.Trim(string(byteBody), "\n")

	assert.Equal(t, "Working!", message)
}
