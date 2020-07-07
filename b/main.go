package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "b")
	fmt.Printf("%#v\n", r.RemoteAddr)
	bss, _ := json.MarshalIndent(r.Header, "", "	")
	fmt.Printf("%v\n", string(bss))
}

func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8022", nil)
}
