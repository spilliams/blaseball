package pkg

type Coded interface {
	error
	StatusCode() int
}

type CodedError struct {
	err  error
	code int
}

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
