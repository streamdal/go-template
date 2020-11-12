package blunder

import "fmt"

type ServiceErr struct {
	Reason string `json:"reason"`
	Err    error  `json:"-"`
}

func (e ServiceErr) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("database error: %s: %v", e.Reason, e.Err)
	}
	return fmt.Sprintf("database error: %s", e.Reason)
}

func (e ServiceErr) Unwrap() error {
	return e.Err
}
