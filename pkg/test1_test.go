package main

import (
	"net/http"
	"os"
	"testing"
)

func TestHealth(t *testing.T) {
	url := os.Getenv("ENDPOINT_URL")
	if len(url) == 0 {
		t.Error("no ENDPOINT_URL")
		t.FailNow()
	}

	res, err := http.Get(url)
	if err != nil {
		t.Errorf("error making http request: %s\n", err)
	}
	if res.StatusCode != 200 {
		t.Errorf("error making http request, non 200 OK status code: %d\n", res.StatusCode)
	}

	t.Log("client: got response!\n")
	t.Logf("client: status code: %d\n", res.StatusCode)
}
