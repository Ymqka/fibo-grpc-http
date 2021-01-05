package http

import (
	"encoding/json"
	"log"
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

func parseFiboStartStop(keys url.Values) (fibo.Params, error) {
	rawStart := keys.Get("start")
	rawStop := keys.Get("stop")
	rawForce := keys.Get("force")

	if rawStart == "" || rawStop == "" {
		return fibo.Params{}, new(emptyStartStop)
	}

	var force bool
	if rawForce == "" {
		force = false
	} else if rawForce == "1" {
		force = true
	}

	start, _ := strconv.ParseUint(rawStart, 10, 32)
	stop, _ := strconv.ParseUint(rawStop, 10, 32)

	fp := fibo.Params{
		Start: uint32(start),
		Stop:  uint32(stop),
		Force: force,
	}

	return fp, nil
}

// FiboPage handles requests
func (server *FiboHTTPServer) FiboPage(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	fiboParams, err := parseFiboStartStop(keys)
	if err != nil {
		log.Printf("failed to parse params, %v", err)
		http.Error(w, err.Error(), 400)
		return
	}

	fiboSeries, err := server.Fibo.FiboRange(fiboParams)
	if err != nil {
		log.Printf("failed to get sequence, %v", err)
		http.Error(w, err.Error(), 500)
		return
	}

	sequence, err := json.Marshal(fiboSeries)
	if err != nil {
		log.Printf("failed to marshal json, %v", err)
		http.Error(w, err.Error(), 500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(sequence)

	return
}

// FiboPageNoCache handle request for fibonacci sequence without cache
func (server *FiboHTTPServer) FiboPageNoCache(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	fiboParams, err := parseFiboStartStop(keys)
	if err != nil {
		log.Printf("failed to parse params, %v", err)
		http.Error(w, err.Error(), 400)
	}

	fiboSeries, err := server.Fibo.FiboRangeNoCache(fiboParams)
	if err != nil {
		log.Printf("failed to get sequence, %v", err)
		http.Error(w, err.Error(), 500)
	}

	sequence, err := json.Marshal(fiboSeries)
	if err != nil {
		log.Printf("failed to marshal json, %v", err)
		http.Error(w, err.Error(), 500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(sequence)

	return
}

func handlers(redisAddr string) http.Handler {
	FiboServer := FiboHTTPServer{Fibo: fibo.Fibonacci{Cache: caching.NewCacheConnection(redisAddr)}}

	r := http.NewServeMux()

	r.HandleFunc("/fibonacci", FiboServer.FiboPage)
	r.HandleFunc("/fibonaccinocache", FiboServer.FiboPageNoCache)

	return r
}

// ServeFiboHTTP connection
func ServeFiboHTTP(addr, redisAddr string) {

	handlers := handlers(redisAddr)

	err := http.ListenAndServe(addr, logRequest(handlers))
	if err != nil {
		log.Fatal(err)
	}

	return
}
