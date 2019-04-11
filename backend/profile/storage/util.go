package storage

import (
	"github.com/google/uuid"
	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

// Dummy is a dummy implementation of Register, managed in-memory.
// Not to be used in Production.
type Dummy struct {
	profiles  map[string]*profilepb.Profile // maps UUID to a user object
	emailToID map[string]string             // maps email to uuid
}

// GetPassword retrieves a password, given a uuid
func GetPassword(userUUID uuid.UUID) {

}
