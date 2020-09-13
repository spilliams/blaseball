package pkg

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
func NewCodedError(e error, c int) CodedError {
	return CodedError{
		err:  e,
		code: c,
	}
}

func (ce CodedError) Error() string {
	return ce.err.Error()
}

func (ce CodedError) StatusCode() int {
	return ce.code
}
