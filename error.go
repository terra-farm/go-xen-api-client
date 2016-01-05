package xenAPI

import (
	"fmt"
)

// Error represents errors returned on xmlrpc request.
type Error struct {
	code    string
	objtype string
	uuid    string
}

// Error() method implements Error interface
func (e *Error) Error() string {
	return fmt.Sprintf("API Error: %s %s %s", e.code, e.objtype, e.uuid)
}

// Code ...
func (e *Error) Code() string {
	return e.code
}

// Type ...
func (e *Error) Type() string {
	return e.objtype
}

// UUID ...
func (e *Error) UUID() string {
	return e.uuid
}
