package session

import (
	"net/http"
	"time"
)

//Start validates if there is an active session, and if not starts a new one
//Returns session id and error
func (sesh *Session) Start(uuid string, username string, ipAddress string, rememberMe bool, w *http.ResponseWriter, r *http.Request) (string, error) {

	session, _ := sesh.cookie.Get(r, sesh.sessionKey)
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

	sessID, _, err := sesh.strg.FindSession(r.Context(), uuid, ipAddress)
	if err == sesh.strg.ErrSessionDoesNotExist() {
		sessID, err = sesh.strg.NewSession(r.Context(), uuid, username, ipAddress, time.Now().Format(sesh.timeFormat), cookie.Expires.Format(sesh.timeFormat))
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	} else {
		err = sesh.strg.SetActive(r.Context(), sessID, time.Now().Format(sesh.timeFormat), cookie.Expires.Format(sesh.timeFormat))
	}

	return sessID, err
}
