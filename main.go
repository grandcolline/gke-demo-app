package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func main() {
	flag.Parse()
	if flag.Arg(0) == "hc" {
		fmt.Println("ok")
	} else {
		fmt.Println("--- Server Start")
		r := chi.NewRouter()
		r.Get("/{version}/", healthHandler)
		r.Get("/{version}/fibo", fiboHandler)
		r.Get("/{version}/down", downHandler)
		http.ListenAndServe(":8080", r)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- versionHandler")
	fmt.Fprint(w, "verion: v1")
}

func fiboHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- fiboHandler")
	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		fmt.Fprint(w, "ERROR")
		return
	}
	fmt.Fprint(w, strconv.Itoa(n)+"番目のフィボナッチ数は、"+strconv.Itoa(fibo(n)))
}

func downHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- downHandler")
	log.Fatal("DOWN!!!")
}

func fibo(n int) int {
	if n < 2 {
		return 1
	}
	return fibo(n-2) + fibo(n-1)
}
