package blunder

import "fmt"

// Validation error is a wrapper on top of ozzo-validation
type ValidationErr struct {
	Field  string      `json:"field"`
	Value  interface{} `json:"value"`
	Reason string      `json:"reason"`
	Err    error       `json:"-"`
}

func (e ValidationErr) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s %s", e.Field, e.Reason)
	}
	return fmt.Sprintf("%s is invalid", e.Field)
}

func (e ValidationErr) Unwrap() error {
	return e.Err
}
