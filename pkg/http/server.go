package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Ymqka/fibo-grpc-http/pkg/caching"

	fibo "github.com/Ymqka/fibo-grpc-http/pkg/fibonacci"
)

// FiboHTTPServer q
type FiboHTTPServer struct {
	Fibo fibo.Fibonacci
}

// FiboPage handles requests
func (server *FiboHTTPServer) FiboPage(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	rawStart := keys.Get("start")
	rawStop := keys.Get("stop")

	if rawStart == "" || rawStop == "" {
		http.Error(w, "start or stop is not provided", 400)
		return
	}

	start, _ := strconv.ParseUint(rawStart, 10, 32)
	stop, _ := strconv.ParseUint(rawStop, 10, 32)

	fiboSeries, err := server.Fibo.Fiborange(uint32(start), uint32(stop))
	if err != nil {
		log.Fatalf("failed to get fibonacci sequence: %v", err)
	}

	fmt.Fprintln(w, fiboSeries)

	return
}

// ServeFiboHTTP connection
func ServeFiboHTTP() {

	FiboServer := FiboHTTPServer{Fibo: fibo.Fibonacci{Cache: caching.NewCacheConnection()}}

	http.HandleFunc("/", FiboServer.FiboPage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
