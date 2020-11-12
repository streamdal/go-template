package blunder

import "fmt"

type DBErr struct {
	Reason string `json:"reason"`
	Err    error  `json:"-"`
}

func (e DBErr) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("database error: %s: %v", e.Reason, e.Err)
	}
	return fmt.Sprintf("database error: %s", e.Reason)
}

func (e DBErr) Unwrap() error {
	return e.Err
}
