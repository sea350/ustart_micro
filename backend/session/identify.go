package session

import (
	"net/http"
	"time"
)

//Identify attempts to pull the userID from the net request session
//Returns session userid and error
func (sesh *Session) Identify(ipAddress string, w *http.ResponseWriter, r *http.Request) (string, error) {

	//TODO:
	// Session called session_please is retreived if it exists
	session, _ := sesh.cookie.Get(r, "session_please")
	// check if uuid exists within the session note: there is inconsistency with checking docid/username.
	uuid, _ := session.Values["UUID"]
	if uuid != nil { // if logged in
		sessID, expiration, err := sesh.strg.FindSession(r.Context(), uuid.(string), ipAddress)
		if err != nil {
			return "", err
		}
		exp, err := time.Parse(expiration, sesh.timeFormat)
		if err != nil {
			return "", err
		}
		if exp.Before(time.Now()) {
			return "", errInvalidSession
		}
		return sessID, nil
	}
	//logic accordingly
	//update active_since in database

	return "", errInvalidSession
}
