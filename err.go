package err

// Error represents a Golang error which has a "class" as well as a "type". The class and
// type attributes of this data type, are meant to help facilitate error handling.
type Error struct {
	errString string
	errClass int
	errType int
	errCause error
}

// New () creates a new data of this type. Always use this function, when you want to
// create new errors.
//
// Input 0: The error (a string describing it).
//
// Input 1: The error's class.
//
// Input 2: The error's type.
//
// 	someErr := err.New ("File could not be opened.", 1, 1, errors.New ("Error foo."))
func New (errString string, errClass, errType int, errCause ... error) (*Error) {
	var cause error

	if len (errCause) > 0 {
		cause = errCause [0]
	} else {
		cause = nil
	}
	return &Error {errString, errClass, errType, cause}
}

// Error () returns the description of the error.
func (err *Error) Error () (string) {
	return err.errString
}

// Class () returns the class of the error.
func (err *Error) Class () (int) {
	return err.errClass
}

// Type () returns the type of the error.
func (err *Error) Type () (int) {
	return err.errType
}

// Unwrap () returns the cause of the error.
func (err *Error) Unwrap () (error) {
	return err.errCause
}
