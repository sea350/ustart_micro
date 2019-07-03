package session

import (
	"github.com/gorilla/sessions"
	"github.com/sea350/ustart_micro/backend/session/storage"
)

//Session is an implementation of the session manager
type Session struct {
	strg   storage.Storage
	cookie *sessions.CookieStore
}

// New returns a new SQL session service
func New(cfg *Config) (*Session, error) {
	// if cfg.useDummy
	strg, err := storage.NewSQL(cfg.StorageConfig)

	sesh := &Session{
		strg:   strg,
		cookie: sessions.NewCookieStore([]byte(cfg.SessionKey)),
	}

	return sesh, err
}
