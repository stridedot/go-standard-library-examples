package expvar_test

import (
	"expvar"
	_ "expvar"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

var (
	ok          = expvar.NewInt("200") // 200 计算器
	notFound    = expvar.NewInt("404") // 404 计数器
	serverError = expvar.NewInt("500") // 500 计数器
)

func TestName(t *testing.T) {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/random", random)
	http.ListenAndServe(":8080", nil)
}

func random(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3)
	var code int

	switch n {
	case 0:
		code = http.StatusOK
		ok.Add(1) // 增加 200 计数器
	case 1:
		code = http.StatusNotFound
		notFound.Add(1) // 增加 404 计数器
	case 2:
		code = http.StatusInternalServerError
		serverError.Add(1) // 增加 500 计数器
	}

	w.WriteHeader(code)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	if err != nil {
		log.Fatal(err)
	}

	ok.Add(1) // 增加 200 计数器
}
