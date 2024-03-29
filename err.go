package err

import (
	"math/big"
)

// Error represents a Golang error which has a "class" as well as a "type". The class and
// type attributes of this data type, are meant to help facilitate error handling in a
// Golang software.
type Error struct {
	errString string
	errClass *big.Int
	errType *big.Int
	secondary error
}

// New () creates a new data of this type. Always use this function, when you want to
// create new errors.
//
// Input 0: The error -- a string describing it.
//
// Input 1: The error's class.
//
// Input 2: The error's type. This could be used to further identify errors falling under
// a common class.
//
// Input 3: If the error you are creating is caused by another error, this value could be
// that error which caused the error you are creating. For example, if you want to create
// error A (File could not be opened). And error A was caued by error B (File does not
// exist). Then the value of this input could be error B. If error A is not caused by any
// error, you should skip this input.
//
// 	someErr := err.New ("File could not be opened", big.NewInt (36), big.NewInt (8),
//		errors.New ("File does not exist"))
func New (errString string, errClass, errType *big.Int, secondary ... error) (*Error) {
	var sec error

	if len (secondary) > 0 {
		sec = secondary [0]
	} else {
		sec = nil
	}
	return &Error {errString, errClass, errType, sec}
}

// Error () returns the description of the error.
func (e *Error) Error () (string) {
	return e.errString
}

// Class () returns the class of the error.
func (e *Error) Class () (*big.Int) {
	// The actual value of the error's class is not returned, to avoid it getting
	// modified outside this package. Rather, a clone of the actual value of the
	// error's class is returned.

	output := big.NewInt (0)
	return output.Add (output, e.errClass)
}

// Type () returns the type of the error.
func (e *Error) Type () (*big.Int) {
	// The actual value of the error's type is not returned, to avoid it getting
	// modified outside this package. Rather, a clone of the actual value of the
	// error's type is returned.

	output := big.NewInt (0)
	return output.Add (output, e.errType)
}

// Unwrap () returns the secondary of the error.
func (e *Error) Unwrap () (error) {
	return e.secondary
}
