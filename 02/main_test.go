package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandlerGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		RootHandler(w, r, HealthHandler)
	})

	reader := strings.NewReader("")
	r, _ := http.NewRequest(http.MethodGet, "/healthz", reader)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response code is %v", resp.StatusCode)
	}
}

func TestHealthHandlerPost(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		RootHandler(w, r, HealthHandler)
	})

	reader := strings.NewReader("")
	r, _ := http.NewRequest(http.MethodPost, "/healthz", reader)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusForbidden {
		t.Errorf("Response code is %v", resp.StatusCode)
	}
}

func TestIndexHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RootHandler(w, r, IndexHandler)
	})

	str := "hello world"
	reader := strings.NewReader(str)
	r, _ := http.NewRequest(http.MethodPost, "/", reader)

	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if string(body) != str {
		t.Errorf("Response err: %s", body)
	}
}
