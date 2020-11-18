// Package Helloworld API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handlers.Hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
