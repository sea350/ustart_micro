package elastic

import "github.com/sea350/ustart_mono/backend/auth/storage"

func (estor *Elastic) Authenticate() error {
	// we figured out a user DNE
	return storage.ErrUserDoesNotExist
}
