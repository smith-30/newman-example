package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "b")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello_b")
}

func main() {
	fmt.Printf("%#v\n", "start service-b")
	http.HandleFunc("/", root)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8022", nil)
}
