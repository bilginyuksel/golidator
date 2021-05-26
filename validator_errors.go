package gorify

import (
	"fmt"
	"strings"
)

var errorMappings = map[string]error{
	"int-between": newGorifyErr("int", "between", "the number: %v, should between min: %v, max: %v"),
	"int-min":     newGorifyErr("int", "min", "the number: %v, should be greater than min: %v"),
	"int-max":     newGorifyErr("int", "max", "the number: %v, should be lower than max: %v"),

	"int64-between": newGorifyErr("int64", "between", "the number: %v, should between min: %v, max: %v"),
	"int64-min":     newGorifyErr("int64", "min", "the number: %v, should be greater than min: %v"),
	"int64-max":     newGorifyErr("int64", "max", "the number: %v, should be lower than max: %v"),

	"string-blank":    newGorifyErr("string", "blank", ""),
	"string-pattern":  newGorifyErr("string", "pattern", ""),
	"string-email":    newGorifyErr("string", "email", ""),
	"string-contains": newGorifyErr("string", "contains", ""),
	"string-between":  newGorifyErr("string", "between", ""),
	"string-min":      newGorifyErr("string", "min", ""),
	"string-max":      newGorifyErr("string", "max", ""),
	"string-size":     newGorifyErr("string", "max", ""),

	"time-between": newGorifyErr("time", "between", ""),
	"time-after":   newGorifyErr("time", "after", ""),
	"time-before":  newGorifyErr("time", "before", ""),
}

type GorifyErr struct {
	Type       string
	Constraint string
	Desc       string
}

func (err *GorifyErr) Error() string {
	return fmt.Sprintf("Type: %v, Constraint: %v, Desc: %v", err.Type, err.Constraint, err.Desc)
}

func (err *GorifyErr) objects(obj ...interface{}) *GorifyErr {
	err.Desc = fmt.Sprintf(err.Desc, obj...)
	return err
}

func newGorifyErr(_type, constraint, desc string, obj ...interface{}) *GorifyErr {
	return &GorifyErr{
		Type:       _type,
		Constraint: constraint,
		Desc:       fmt.Sprintf(desc, obj...),
	}
}

// UpdateErrResponse you can override validation error behavior using this method.
// For example if you want to use the validation err response in your business logic
// and if you have a special error format you can configure it and then whenever an error
// occurs at validation it will throw the error you have configured.
func UpdateErrResponse(_type, constraint string, response error) {
	_type = strings.ToLower(_type)
	constraint = strings.ToLower(constraint)

	scenario := fmt.Sprintf("%s-%s", _type, constraint)
	errorMappings[scenario] = response
}
