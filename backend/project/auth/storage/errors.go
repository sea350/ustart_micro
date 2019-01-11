package store

import "errors"

var (
	//ErrNoResults No results found
	ErrNoResults = errors.New("No results found")

	//ErrTooManyResults Too many results found (>1 results)
	ErrTooManyResults = errors.New("Too many results, a crititcal error has occurred")
)
