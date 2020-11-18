package handler

import (
	"encoding/json"
	"io"
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
func Hello(w http.ResponseWriter, r *http.Request) {
	resp := ResponseMessage{Message: "hello world"}
	res_json, _ := json.Marshal(resp)
	io.WriteString(w, string(res_json))
}

// HelloResponse is an response with helloworld message.
//
// swagger:response HelloResponse
type ResponseMessage struct {
	Message string
}
