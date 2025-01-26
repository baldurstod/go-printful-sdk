package printfulsdk

import "fmt"

type HTTPError struct {
	err     error
	context any
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%s:%v", e.err.Error(), e.context)
}

func (e *HTTPError) Unwrap() error {
	return e.err
}

func NewHTTPError(err error, context any) error {
	return &HTTPError{
		err:     err,
		context: context,
	}
}
