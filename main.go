package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func fiboPage(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	rawStart := keys.Get("start")
	rawStop := keys.Get("stop")

	if rawStart == "" || rawStop == "" {
		http.Error(w, "start or stop is not provided", 400)
		return
	}

	start, _ := strconv.Atoi(rawStart)
	stop, _ := strconv.Atoi(rawStop)

	fiboSeries := fibonacci(start, stop)

	fmt.Fprintln(w, fiboSeries)

	return
}

func handleRequests() {
	http.HandleFunc("/fibo", fiboPage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
