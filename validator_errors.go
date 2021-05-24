package gorify

import "fmt"

type GorifyErr struct {
	Type       string
	Constraint string
	Message    string
}

func (err *GorifyErr) Error() string {
	return fmt.Sprintf("Type: %v, Constraint: %v, Message: %v", err.Type, err.Constraint, err.Message)
}

func newGorifyErr(_type, constraint, message string, messageObjects ...interface{}) *GorifyErr {
	msg := fmt.Sprintf(message, messageObjects...)
	return &GorifyErr{
		Type:       _type,
		Constraint: constraint,
		Message:    msg,
	}
}
