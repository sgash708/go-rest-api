package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestMain_成功 httpアクセス成功
func TestMain_成功(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(Hello))
	defer testServer.Close()

	res, err := http.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}
	hello, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error("response is not 200")
	}
	if string(hello) != "Hello from handler!\n" {
		t.Error("do not match")
	}
}
