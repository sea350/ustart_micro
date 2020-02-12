package project

import (
	"fmt"
	"net/http"
)

//Ping is used to test a connection
func (rapi *RESTAPI) Ping(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w, "Pong")
}
