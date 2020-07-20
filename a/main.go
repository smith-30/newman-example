package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/smith-30/newman-example/rdb"
)

type User struct {
	ID   int
	Name string
}

func NewUser() User {
	return User{
		ID:   1,
		Name: "user",
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "a")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(NewUser())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func showUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] != "1" {
		http.Error(w, "invalid user id", http.StatusBadGateway)
		return
	}

	res, err := json.Marshal(NewUser())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	_, err := rdb.Mysql("test", "test_pass_hoge", "mysql", "3306", "testdb", true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", "start service-a")

	r := mux.NewRouter()
	r.HandleFunc("/", root)
	r.HandleFunc("/user", createUser).Methods("POST")
	r.HandleFunc("/user/{id}", showUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8011", r))
}
