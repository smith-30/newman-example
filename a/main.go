package main

import (
	"fmt"
	"net/http"

	"github.com/smith-30/cidind/rdb"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "a")
}

func main() {
	_, err := rdb.Mysql("test", "test_pass_hoge", "mysql", "3306", "testdb", true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", "ok")
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe("0.0.0.0:8011", nil)
}
