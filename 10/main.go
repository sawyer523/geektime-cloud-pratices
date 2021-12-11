package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"syscall"
	"time"

	"github.com/geektime-cloud-pratices/10/metrics"

	"gopkg.in/yaml.v2"

	"github.com/fsnotify/fsnotify"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func randInt(min int, max int) time.Duration {
	rand.Seed(time.Now().UTC().UnixNano())
	return time.Duration(min + rand.Intn(max-min))
}

func RootHandler(w http.ResponseWriter, r *http.Request, f func(http.ResponseWriter, *http.Request)) {
	for k, v := range r.Header {
		for _, val := range v {
			w.Header().Add(k, val)
		}
	}
	fnName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	metric := metrics.New()

	w.Header().Set("version", os.Getenv("VERSION"))
	timer := time.NewTicker(randInt(0, 2000) * time.Millisecond)
	<-timer.C
	f(w, r)
	code := reflect.ValueOf(w).Elem().FieldByName("status")
	fmt.Println(fmt.Sprint(code))
	defer func() {
		metric.ObserveTotal(fnName, r.Method)
		metric.Count(fnName, r.Method, fmt.Sprint(code))
		r.Body.Close()
	}()
	logrus.Debug("ip: ", r.RemoteAddr, ", http code: ", code)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(403)
		io.WriteString(w, "forbidden")
		return
	}
	w.WriteHeader(http.StatusOK)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	w.Write(body)
	w.WriteHeader(http.StatusOK)
}

type Bootstrap struct {
	LogLevel string `yaml:"logLevel"`
	Port     string `yaml:"port"`
}

var done chan struct{}

var svc http.Server
var conf string

func init() {
	flag.StringVar(&conf, "conf", "./", "config path")
}

func main() {
	flag.Parse()
	metrics.Register()
	logrus.Info("http starting")
	os.Setenv("VERSION", "3")

	done = make(chan struct{})
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		<-c
		stop()
	}()
	flconf := path.Join(conf, "config")
	go Watch(flconf, done)

	bs, err := readConfig(flconf)
	if err != nil {
		logrus.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		RootHandler(w, r, HealthHandler)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RootHandler(w, r, IndexHandler)
	})

	mux.Handle("/metrics", promhttp.Handler())

	svc = http.Server{
		Addr:    fmt.Sprintf(":%s", bs.Port),
		Handler: mux,
	}

	if err := svc.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logrus.Fatal(err)
	} else {
		logrus.Info("httpserver stopped")
	}

}

func stop() {
	logrus.Info("stopping")
	done <- struct{}{}
	svc.Shutdown(context.Background())
}

func Watch(p string, done <-chan struct{}) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logrus.WithField("watch", "config").Fatal(err)
	}
	defer watcher.Close()
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				logrus.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					logrus.Println("modified file:", event.Name)
					name := filepath.Base(event.Name)
					if name == "config" {
						readConfig(event.Name)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(p)
	if err != nil {
		logrus.WithField("watcher", "add").WithField("file", "config").Fatal(err)
	}
	<-done
}

func readConfig(name string) (Bootstrap, error) {
	var c Bootstrap
	data, err := ioutil.ReadFile(name)
	if err != nil {
		logrus.WithField("ioutil", "read").Errorln(err)
		return c, fmt.Errorf("ioutil.ReadFile: %v", err)
	}
	if err := yaml.Unmarshal(data, &c); err != nil {
		logrus.WithField("modify", "invalid").Errorln(err)
		return c, fmt.Errorf("yaml.Unmarshal: %v", err)
	}
	level, err := logrus.ParseLevel(c.LogLevel)
	if err != nil {
		logrus.Info("logrus ParseLevel Error: ", c.LogLevel)
		return c, fmt.Errorf("logrus.ParseLevel: %v", err)
	}
	logrus.SetLevel(level)
	return c, nil
}
