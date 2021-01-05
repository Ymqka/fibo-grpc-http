package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	fibo "github.com/Ymqka/fibo-grpc-http/pkg/fibonacci"
)

func TestRouting_FiboPage(t *testing.T) {
	srv := httptest.NewServer(handlers(":6379"))
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/fibonacci?start=0&stop=3", srv.URL))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf(res.Status)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	var fiboSeq []fibo.FibonacciSequence
	json.Unmarshal(body, &fiboSeq)

	fiboSeqExpect := []fibo.FibonacciSequence{
		{ID: 0, Num: big.NewInt(0)},
		{ID: 1, Num: big.NewInt(1)},
		{ID: 2, Num: big.NewInt(1)},
		{ID: 3, Num: big.NewInt(2)},
	}

	if !reflect.DeepEqual(fiboSeq, fiboSeqExpect) {
		t.Errorf("got %v, wanted %v", fiboSeq, fiboSeqExpect)
	}
}

func TestRouting_FiboPageNoCache(t *testing.T) {
	srv := httptest.NewServer(handlers(":6379"))
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/fibonaccinocache?start=0&stop=3", srv.URL))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf(res.Status)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	var fiboSeq []fibo.FibonacciSequence
	json.Unmarshal(body, &fiboSeq)

	fiboSeqExpect := []fibo.FibonacciSequence{
		{ID: 0, Num: big.NewInt(0)},
		{ID: 1, Num: big.NewInt(1)},
		{ID: 2, Num: big.NewInt(1)},
		{ID: 3, Num: big.NewInt(2)},
	}

	if !reflect.DeepEqual(fiboSeq, fiboSeqExpect) {
		t.Errorf("got %v, wanted %v", fiboSeq, fiboSeqExpect)
	}
}

func TestRouting_FiboError_EmptyStartStop(t *testing.T) {
	srv := httptest.NewServer(handlers(":6379"))
	defer srv.Close()

	res, _ := http.Get(fmt.Sprintf("%s/fibonacci?start=0", srv.URL))
	if res.StatusCode != 400 {
		t.Error("response code for empty params is not 400")
	}
}

func TestRouting_FiboError_ForceFlag(t *testing.T) {
	srv := httptest.NewServer(handlers(":6379"))
	defer srv.Close()

	res, _ := http.Get(fmt.Sprintf("%s/fibonacci?start=0&stop=999999", srv.URL))
	if res.StatusCode == http.StatusOK {
		t.Error("stop higher than 10000 without force flag returned 200 code ")
	}
}
