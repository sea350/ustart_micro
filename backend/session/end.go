package session

import (
	"net/http"
)

//End kills a session, meant for logout
//Returns error
func (sesh *Session) End(w *http.ResponseWriter, r *http.Request) error {

	// Session called session_please is retreived if it exists
	session, _ := sesh.cookie.Get(r, "session_please")
	// check if uuid exists within the session note: there is inconsistency with checking docid/username.
	uuid, _ := session.Values["UUID"]
	if uuid != nil { // if logged in
		session.Options.MaxAge = -1 // kills session
		session.Save(r, *w)         // save changes to session
		return nil
	}

	//TODO:
	//Double check edge cases
	//have a proper failure return
	return nil
}
