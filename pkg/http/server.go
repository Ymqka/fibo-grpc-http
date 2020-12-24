package http

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Ymqka/fibo-grpc-http/pkg/caching"

	fibo "github.com/Ymqka/fibo-grpc-http/pkg/fibonacci"
)

// FiboHTTPServer q
type FiboHTTPServer struct {
	Fibo fibo.Fibonacci
}

type emptyStartStop struct {
}

func (e *emptyStartStop) Error() string { return "empty start or stop" }

func parseFiboStartStop(keys url.Values) (uint32, uint32, error) {
	rawStart := keys.Get("start")
	rawStop := keys.Get("stop")

	if rawStart == "" || rawStop == "" {
		return 0, 0, &emptyStartStop{}
	}

	start, _ := strconv.ParseUint(rawStart, 10, 32)
	stop, _ := strconv.ParseUint(rawStop, 10, 32)

	return uint32(start), uint32(stop), nil
}

// FiboPage handles requests
func (server *FiboHTTPServer) FiboPage(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	start, stop, err := parseFiboStartStop(keys)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	fiboSeries, err := server.Fibo.Fiborange(start, stop)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	fmt.Fprintln(w, fiboSeries)

	return
}

// FiboPageNoCache handle request for fibonacci sequence without cache
func (server *FiboHTTPServer) FiboPageNoCache(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	start, stop, err := parseFiboStartStop(keys)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	fiboSeries, err := server.Fibo.FiborangeNoCache(start, stop)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	fmt.Fprintln(w, fiboSeries)

	return
}

// ServeFiboHTTP connection
func ServeFiboHTTP() {

	FiboServer := FiboHTTPServer{Fibo: fibo.Fibonacci{Cache: caching.NewCacheConnection("redis:6379")}}

	http.HandleFunc("/fibonacci", FiboServer.FiboPage)
	http.HandleFunc("/fibonaccinocache", FiboServer.FiboPageNoCache)

	http.ListenAndServe(":10000", nil)
}
