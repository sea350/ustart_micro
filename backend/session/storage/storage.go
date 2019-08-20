package storage

// Storage is a database-agnostic interface for persisting session data
type Storage interface {
	NewSession(string, string, string, string) (string, error) //pass uuid, username, ip, logintime | get sessionID, error
	FindSession(string, string) (string, error)                //pass uuid, ip | get sessionID, error
	SetActive(string, string) error                            //pass sessionID, loginTime | get error
	EndSession(string) error                                   //pass sessionID | get error
}
