package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func main() {

	http.HandleFunc("/token", handler)

	fmt.Println("Starting the http server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
