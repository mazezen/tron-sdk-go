package api

import "fmt"

type SendRequestError struct {
	Err error
}

func (err *SendRequestError) Error() string {
	return fmt.Sprintf("%v", err.Err)
}
