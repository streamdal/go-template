package blunder

import "fmt"

// Any errors that are coming from the standard library
type StandardErr struct {
	Reason string `json:"message"`
	Err    error  `json:"-"`
}

func (e StandardErr) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("error: %s: %v", e.Reason, e.Err)
	}
	return fmt.Sprintf("error: %s", e.Reason)
}

func (e StandardErr) Unwrap() error {
	return e.Err
}
