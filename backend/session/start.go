package session

import (
	"net/http"
)

//Start validates if there is an active session, and if not starts a new one
//Returns session id and error
func (sesh *Session) Start(uuid string, ipAddress string, w *http.ResponseWriter, r *http.Request) (string, error) {

	//TODO:
	//Self.identity
	//Start new session
	//Update database

	return "", nil
}
