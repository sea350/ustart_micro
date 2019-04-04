package storage

import (
	"github.com/google/uuid"
	"github.com/sea350/ustart_micro/backend/auth/authpb"
)

// Dummy is a dummy implementation of Register, managed in-memory.
// Not to be used in Production.
type Dummy struct {
	users     map[string]*authpb.User // maps UUID to a user object
	emailToID map[string]string       // maps email to uuid
}

// GetPassword retrieves a password, given a uuid
func GetPassword(userUUID uuid.UUID) {

}
