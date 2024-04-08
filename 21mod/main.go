package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", ServeWeb).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hello from mod")
}

func ServeWeb(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the world of GoLang</h1>"))
}
