package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"

	"github.com/sirupsen/logrus"
)

func RootHandler(w http.ResponseWriter, r *http.Request, f func(http.ResponseWriter, *http.Request)) {
	defer r.Body.Close()
	for k, v := range r.Header {
		for _, val := range v {
			w.Header().Add(k, val)
		}
	}
	w.Header().Set("version", os.Getenv("VERSION"))
	f(w, r)
	code := reflect.ValueOf(w).Elem().FieldByName("status")
	logrus.Info("ip: ", r.RemoteAddr, ", http code: ", code)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(403)
		io.WriteString(w, "forbidden")
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	w.Write(body)
	fmt.Println("body: ", string(body))
}

func main() {
	logrus.Info("http starting")
	os.Setenv("VERSION", "3")
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		RootHandler(w, r, HealthHandler)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RootHandler(w, r, IndexHandler)
	})

	logrus.Fatal(http.ListenAndServe(":8080", mux))
}
