package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/abhi-alt/mongodb-connect/routers"
)

func main() {
	fmt.Println("MongoDB Demo")
	fmt.Println("Your connection has started....")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
}
