package blunder

import "fmt"

type AuthErr struct {
	Reason string `json:"reason"`
	Err    error  `json:"-"`
}

func (e AuthErr) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("database error: %s: %v", e.Reason, e.Err)
	}
	return fmt.Sprintf("database error: %s", e.Reason)
}

func (e AuthErr) Unwrap() error {
	return e.Err
}
