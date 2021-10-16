package server_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheckEndPoint(t *testing.T) {
	t.Parallel()

	port := 3000

	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/healthcheck", port), nil)

	client := http.Client{}
	response, _ := client.Do(request)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	byteBody, _ := ioutil.ReadAll(response.Body)
	message := strings.Trim(string(byteBody), "\n")

	assert.Equal(t, "Working!", message)

	response.Body.Close()
}
