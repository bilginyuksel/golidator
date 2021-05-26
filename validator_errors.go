package gorify

import (
	"fmt"
	"strings"
)

var errorMappings = map[string]error{
	"int-between": newGorifyErr("int", "between", "given number: {}, should between min: {}, max: {}"),
	"int-min":     newGorifyErr("int", "min", "given number: {}, should be greater than min: {}"),
	"int-max":     newGorifyErr("int", "max", "given number: {}, should be lower than max: {}"),

	"int64-between": newGorifyErr("int64", "between", "given number: {}, should between min: {}, max: {}"),
	"int64-min":     newGorifyErr("int64", "min", "given number: {}, should be greater than min: {}"),
	"int64-max":     newGorifyErr("int64", "max", "given number: {}, should be lower than max: {}"),

	"string-blank":    newGorifyErr("string", "blank", "given string should not be blank"),
	"string-pattern":  newGorifyErr("string", "pattern", "given string did not match with the regexp pattern"),
	"string-email":    newGorifyErr("string", "email", "given string is not an email"),
	"string-contains": newGorifyErr("string", "contains", "given string {}, does not contain {}"),
	"string-between":  newGorifyErr("string", "between", "given string length should between min: {}, max: {}"),
	"string-min":      newGorifyErr("string", "min", "given string length should be greater than min: {}"),
	"string-max":      newGorifyErr("string", "max", "given string length should be smaller than max: {}"),
	"string-size":     newGorifyErr("string", "size", "given string size should be equal to size: {}"),

	"time-between": newGorifyErr("time", "between", "given time should between {} and {}"),
	"time-after":   newGorifyErr("time", "after", "given time should before {}"),
	"time-before":  newGorifyErr("time", "before", "given time should after {}"),
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
	err.Desc = fmt.Sprintf(strings.Replace(err.Desc, "{}", "%v", -1), obj...)
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
