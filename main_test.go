package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerReturnsGreetingsWithQueryParameter(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8000/a-path?name=lioda", nil) // step 1

	handler(res, req) // step 2

	content, _ := ioutil.ReadAll(res.Body)
	expected := "Hi there, my name is lioda!"
	if string(content) != expected { // step 3
		t.Errorf("Expected %s, got %s", expected, string(content))
	}
}

