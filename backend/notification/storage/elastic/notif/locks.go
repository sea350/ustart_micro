package elasticstore

import (
	"sync"
)

var newUserLock sync.Mutex
