package blunder

import "fmt"

type NotFoundErr struct {
	Reason string `json:"reason"`
	Err    error  `json:"-"`
}

func (e NotFoundErr) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("database error: %s: %v", e.Reason, e.Err)
	}
	return fmt.Sprintf("database error: %s", e.Reason)
}

func (e NotFoundErr) Unwrap() error {
	return e.Err
}
