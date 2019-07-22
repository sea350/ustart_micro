package sqlstore

import "errors"

//Having errors globalized like this makes it easier for higher level packages to identify specific problems
//and implement logic for special error cases
//the variables are not exported so they cant be modified by anything outside the package

var (
	// errImproperImport is when the string is incorrectly inputted
	errImproperImport = errors.New("The string imported is not in base64 format")
)
