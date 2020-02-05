package elasticstore

import (
	"sync"
)

var newUserLock sync.Mutex

var modProjectsLock sync.Mutex
