package session

import (
	"net/http"
	"time"
)

//End kills a session, meant for logout
//Returns error
func (sesh *Session) End(ipAddress string, w *http.ResponseWriter, r *http.Request) error {

	// Session called session_please is retreived if it exists
	session, _ := sesh.cookie.Get(r, sesh.sessionKey)
	// check if uuid exists within the session note: there is inconsistency with checking docid/username.
	uuid, _ := session.Values["UUID"]
	if uuid != nil { // if logged in
		session.Options.MaxAge = -1 // kills session
		session.Save(r, *w)         // save changes to session

		sessID, _, err := sesh.strg.FindSession(r.Context(), uuid.(string), ipAddress)
		if err != nil {
			//PLACEHOLDER
			//session is already closed this doesn't really matter
		} else {
			err := sesh.strg.EndSession(r.Context(), sessID, time.Now().Format(sesh.timeFormat))
			if err != nil {
				//PLACEHOLDER
				//session is already closed this doesn't really matter
			}
		}

		return nil
	}

	//TODO:
	//Double check edge cases

	return nil
}
