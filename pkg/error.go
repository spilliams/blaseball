package pkg

import "fmt"

// Coded represents an error that can be asked for its status code.
type Coded interface {
	error
	StatusCode() int
}

// CodedError represents an error that can be given a status code.
type CodedError struct {
	err  error
	code int
}

// NewCodedError returns a new error with status code
func NewCodedError(c int, e error) CodedError {
	return CodedError{
		err:  e,
		code: c,
	}
}

// NewCodedErrorf returns a new error with status code
func NewCodedErrorf(c int, format string, parts ...interface{}) CodedError {
	return CodedError{
		err:  fmt.Errorf(format, parts...),
		code: c,
	}
}

func (ce CodedError) Error() string {
	return ce.err.Error()
}

func (ce CodedError) StatusCode() int {
	return ce.code
}
