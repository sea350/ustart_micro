package auth

import (
	"sync"
)

var newUserLock sync.Mutex
