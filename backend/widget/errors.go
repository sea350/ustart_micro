package widget

import "errors"

var errStorageDiscrepancy = errors.New("There was a discrepancy in how an item has been stored, this isn't supposed to happen")

//ErrStorageDiscrepancy standardzized package wide error
func ErrStorageDiscrepancy() error { return errStorageDiscrepancy }
