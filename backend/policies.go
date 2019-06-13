package backend

import (
	"net/http"
	"regexp"
)

//setCORS sets CORS policy and checks for proper request method
func setCORS(w *http.ResponseWriter, req *http.Request) bool {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	if req.Method == "OPTIONS" {
		return false
	}

	return true
}

func validUsername(username string) bool {
	rxUname := regexp.MustCompile(`/^[a-zA-Z0-9][\w]*\.?[\w]*[a-zA-Z0-9]+$/`)
	if !rxUname.MatchString(username) {
		return false
	}
	return true
}
