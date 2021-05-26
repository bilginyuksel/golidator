package gorify

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	emailRegExp                  = "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"
	notBlankErr                  = "the string is blank"
	regExpMatchErr               = "value: %v, does not match with the pattern"
	betweenStrLengthIsNotInRange = "given string is not in range between the numbers in tag"
)

var (
	notBlank = func(value string, tags reflect.StructTag) error {
		blank, ok := tags.Lookup("blank")

		if !ok || blank == "true" || strings.TrimSpace(value) != "" {
			return nil
		}

		return errorMappings["string-blank"]
	}

	matchesRegExp = func(value string, tags reflect.StructTag) error {
		pattern, ok := tags.Lookup("pattern")

		if !ok {
			return nil
		}

		regExp, err := regexp.Compile(pattern)
		if err != nil {
			panic(fmt.Errorf("regular expression could not compiled, err: %v", err))
		}

		if !regExp.MatchString(value) {
			return errorMappings["string-pattern"]
		}

		return nil
	}

	emailStringField = func(value string, tags reflect.StructTag) error {
		if _, ok := tags.Lookup("email"); ok {
			regExp, _ := regexp.Compile(emailRegExp)
			if !regExp.MatchString(value) {
				return errorMappings["string-email"]
			}
		}
		return nil
	}

	containsString = func(value string, tags reflect.StructTag) error {
		if containsStr, ok := tags.Lookup("contains"); ok {
			if !strings.Contains(value, containsStr) {
				return errorMappings["string-contains"].(*GorifyErr).objects(value, containsStr)
			}
		}
		return nil
	}

	lengthBetween = func(value string, tags reflect.StructTag) error {
		if lengthBetweenStr, ok := tags.Lookup("between"); ok {
			minMaxSplitted := strings.Split(lengthBetweenStr, "-")
			if len(minMaxSplitted) != 2 {
				panic("min and max values should be separated by '-'")
			}
			min, err := strconv.Atoi(minMaxSplitted[0])
			max, err := strconv.Atoi(minMaxSplitted[1])

			if err != nil {
				panic("min and max values should be integer")
			}

			if min > len(value) || max < len(value) {
				return errorMappings["string-between"].(*GorifyErr).objects(min, max)
			}
		}
		return nil
	}

	minLength = func(value string, tags reflect.StructTag) error {
		if strMinLength, exists := tags.Lookup("min"); exists {
			min, err := strconv.Atoi(strMinLength)

			if err != nil {
				panic("min length should be integer")
			}

			if len(value) < min {
				return errorMappings["string-min"].(*GorifyErr).objects(min)
			}
		}

		return nil
	}

	maxLength = func(value string, tags reflect.StructTag) error {
		if strMaxLength, exists := tags.Lookup("max"); exists {
			max, err := strconv.Atoi(strMaxLength)

			if err != nil {
				panic("max length should be integer")
			}

			if len(value) > max {
				return errorMappings["string-max"].(*GorifyErr).objects(max)
			}
		}
		return nil
	}

	lengthEqual = func(value string, tags reflect.StructTag) error {
		if strSize, exists := tags.Lookup("size"); exists {
			size, err := strconv.Atoi(strSize)

			if err != nil {
				panic("size should be integer")
			}

			if len(value) != size {
				return errorMappings["string-size"].(*GorifyErr).objects(size)
			}
		}
		return nil
	}

	setDefault = func(reflectField reflect.StructField, reflectValue reflect.Value) error {
		tags := reflectField.Tag
		currValue := reflectValue.String()
		defaultValue, exists := tags.Lookup("default")

		// if tag does not exist or currValue is not empty do nothing
		if !exists || currValue != "" {
			return nil
		}

		if reflectValue.CanSet() {
			reflectValue.SetString(defaultValue)
			return nil
		}

		return errors.New("default value could not set")
	}
)
