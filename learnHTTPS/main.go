package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")
	io.WriteString(w, "My website")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / hello request")
	io.WriteString(w, "Hello")
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":42069", nil)
	if err != nil {
		log.Fatal(err)
	}
}
