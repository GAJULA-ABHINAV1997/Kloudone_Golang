package main

import (
	"fmt"
	"log"
	"net/http"

	"Assignment_7/gorilla_mux/router"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
