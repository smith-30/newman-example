package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/smith-30/cidind/rdb"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "b")
	fmt.Printf("%#v\n", r.RemoteAddr)
	bss, _ := json.MarshalIndent(r.Header, "", "	")
	fmt.Printf("%v\n", string(bss))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "b")
	fmt.Printf("%#v\n", r.RemoteAddr)
	bss, _ := json.MarshalIndent(r.Header, "", "	")
	fmt.Printf("%v\n", string(bss))
}

func main() {
	_, err := rdb.Mysql("test", "test_pass_hoge", "mysql", "3306", "testdb", true)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", root)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe("0.0.0.0:8011", nil)
}
