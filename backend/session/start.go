package session

import (
	"net/http"
	"time"
)

//Start validates if there is an active session, and if not starts a new one
//Returns session id and error
func (sesh *Session) Start(uuid string, username string, ipAddress string, rememberMe bool, w *http.ResponseWriter, r *http.Request) (string, error) {

	session, _ := sesh.cookie.Get(r, "session_please")
	session.Values["UUID"] = uuid
	session.Values["Username"] = username
	cookie := http.Cookie{Name: session.Values["UUID"].(string), Value: "user", Path: "/"}
	if rememberMe {
		cookie.Expires = time.Now().Add(60 * 24 * time.Hour)
	} else {
		session.Options.MaxAge = 0
	}

	http.SetCookie(*w, &cookie)
	session.Save(r, *w)

	//TODO:
	//Update database

	return "", nil
}
