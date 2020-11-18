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
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// swagger:operation GET /hello tag1 tag2 login
//
// Returns all pets from the system that the user has access to
//
// Could be any pet
//
// ---
// produces:
// - application/json
// parameters:
// - name: name
//   in: query
//   description: user name
//   required: true
//   type: string
// - name: password
//   in: query
//   description: user password
//   required: true
//   type: string
// responses:
//   '200':
//     description: login success
//     schema:
//       type: object
//       items:
//         "$ref": "#/responses/HelloResponse"
//   default:
//     description: unexpected error
//     type: object
//     schema:
//       "$ref": "#/responses/HelloResponse"
func hello(w http.ResponseWriter, r *http.Request) {
	resp := ResponseMessage{Message: "hello world"}
	res_json, _ := json.Marshal(resp)
	io.WriteString(w, string(res_json))
}
func main() {
	http.HandleFunc("/hello", hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// HelloResponse is an response with helloworld message.
//
// swagger:response HelloResponse
type ResponseMessage struct {
	Message string
}
