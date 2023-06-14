package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type authentication struct {
	userName string
	password string
}

// populate the username and password from environment vars
var appAuth = authentication{
	userName: os.Getenv("APP_USERNAME"),
	password: os.Getenv("APP_PASSWORD"),
}

func main() {

	http.HandleFunc("/token", handler)
	fmt.Println("Starting the http server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {

	userName, password, ok := request.BasicAuth()

	if !ok {
		fmt.Println("Unable to parse Auth info")
		writer.WriteHeader(401)
		return
	}

	if userName != appAuth.userName || password != appAuth.password {
		fmt.Println("Authentication failed")
		writer.WriteHeader(401)
		return
	}

	fmt.Println("Authentication succeeded")
	fmt.Fprintln(writer, "Welcome", userName)
	writer.WriteHeader(200)
	return
}
