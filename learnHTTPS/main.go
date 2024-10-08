package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")
	fmt.Printf("got / request, first(%v) = %s, second(%v) = %v\n body:%v\n",
		hasFirst, first, hasSecond, second, string(body))
	io.WriteString(w, "My website")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / hello request")
	io.WriteString(w, "Hello")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":42069", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		log.Fatal(err)
	}
}
