package storage

import "context"

// Storage is a database-agnostic interface for persisting session data
type Storage interface {
	NewSession(context.Context, string, string, string, string) (string, error) //pass uuid, username, ip, logintime, expiration | get sessionID, error
	FindSession(context.Context, string, string) (string, error)                //pass uuid, ip | get sessionID, error
	SetActive(context.Context, string, string, string) error                    //pass sessionID, loginTime, expiration | get error
	EndSession(context.Context, string, string) error                           //pass sessionID, expiration | get error
}
