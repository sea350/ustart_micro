package awsstore

import "errors"

var (
	// errImproperImport is when the string is incorrectly inputted
	errImproperImport = errors.New("The string imported is not in base64 format")
)