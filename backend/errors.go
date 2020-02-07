package backend

import "errors"

var errNilResponse = errors.New("One or more microservice responses returned null")

var errProblemLoadingArr = errors.New("One or more items couldn't be loaded in an array")
