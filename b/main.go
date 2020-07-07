package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "b")
	fmt.Printf("%#v\n", r.RemoteAddr)
	bss, _ := json.MarshalIndent(r.Header, "", "	")
	fmt.Printf("%v\n", string(bss))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello_b")
	fmt.Printf("%#v\n", r.RemoteAddr)
	bss, _ := json.MarshalIndent(r.Header, "", "	")
	fmt.Printf("%v\n", string(bss))
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8022", nil)
}
