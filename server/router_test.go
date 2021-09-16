package server_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func Test_EndToEnd_Healthcheck(t *testing.T) {
	port := 3000

	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/healthcheck", port), nil)

	client := http.Client{}
	response, _ := client.Do(request)

	if http.StatusOK != response.StatusCode {
		t.Errorf("status code got %q, want %q", response.StatusCode, http.StatusOK)
	}

	byteBody, _ := ioutil.ReadAll(response.Body)

	want := "Working!"
	got := strings.Trim(string(byteBody), "\n")

	if want != got {
		t.Errorf("body got %q, want %q", got, want)
	}

	response.Body.Close()
}
