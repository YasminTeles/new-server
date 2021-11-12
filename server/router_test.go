//nolint:exhaustivestruct,wsl
package server_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/YasminTeles/new-server/server"
	"github.com/YasminTeles/new-server/settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
}

func (suite *ServerTestSuite) SetupSuite() {
	settings.LoadSettings()
	go server.ListenAndServe()

	time.Sleep(1 * time.Second)
}

func (suite *ServerTestSuite) TestHealthcheckEndPoint() {
	url := fmt.Sprintf("http://localhost:%s/healthcheck", settings.Config.Port)

	request, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)

	client := http.Client{}
	response, _ := client.Do(request)

	assert.Equal(suite.T(), http.StatusOK, response.StatusCode)

	byteBody, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	message := strings.Trim(string(byteBody), "\n")

	assert.Equal(suite.T(), "Working!", message)
}

func TestClientTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ServerTestSuite))
}
