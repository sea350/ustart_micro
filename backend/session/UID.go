package session

import (
	"net/http"
)

//UID returns uuid if there is an active session, error if not
//Returns session userid and error
func (sesh *Session) UID(r *http.Request) (string, error) {

	session, _ := sesh.cookie.Get(r, "session_please")
	uuid, _ := session.Values["UUID"]
	if uuid != nil { // if logged in
		return uuid.(string), nil
	}

	return "", errInvalidSession
}
