package session

import (
	"net/http"
)

//Identify attempts to pull the userID from the net request session
//Returns session userid and error
func (sesh *Session) Identify(w *http.ResponseWriter, r *http.Request) (string, error) {

	//TODO:
	//cookie.Get
	//logic accordingly
	//update active_since in database

	return "", nil
}
