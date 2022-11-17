package gl

import "fmt"

type wrapError struct {
	src error
	msg string
}

func (we *wrapError) Error() string {
	return fmt.Sprint(we.msg, "\n", we.src.Error())
}

func ErrorWrap(err error, msg string) error {
	return &wrapError{src: err, msg: msg}
}
