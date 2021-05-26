package gorify

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

var (
	maxInt64 = func(value int64, tags reflect.StructTag) error {
		if maxStr, ok := tags.Lookup("max"); ok {
			maxInt64, err := parseInt64(maxStr)
			if err != nil {
				panic("max value should be int64")
			}
			if maxInt64 < value {
				return errorMappings["int64-max"].(*GorifyErr).objects(value, maxInt64)
			}
		}
		return nil
	}

	minInt64 = func(value int64, tags reflect.StructTag) error {
		if minStr, ok := tags.Lookup("min"); ok {
			minInt64, err := parseInt64(minStr)
			if err != nil {
				panic("min value should be in64")
			}
			if minInt64 > value {
				return errorMappings["int64-min"].(*GorifyErr).objects(value, minInt64)
			}
		}
		return nil
	}

	betweenInt64 = func(value int64, tags reflect.StructTag) error {
		if betweenStr, ok := tags.Lookup("between"); ok {
			betweenStrSplitted := strings.Split(betweenStr, "-")
			if len(betweenStrSplitted) != 2 {
				panic("min and max should be separated by '-'")
			}
			minInt64, err := parseInt64(betweenStrSplitted[0])
			maxInt64, err := parseInt64(betweenStrSplitted[1])

			if err != nil {
				panic("min and max values should be int64")
			}

			if minInt64 > value || maxInt64 < value {
				return errorMappings["int64-between"].(*GorifyErr).objects(value, minInt64, maxInt64)
			}
		}
		return nil
	}

	setDefaultInt64 = func(reflectField reflect.StructField, reflectValue reflect.Value) error {
		tags := reflectField.Tag
		currValue := reflectValue.Int()
		defaultValue, exists := tags.Lookup("default")

		if !exists || currValue != 0 {
			return nil
		}

		intValue, err := parseInt64(defaultValue)

		if err != nil {
			return errors.New("please write int value to default tag")
		}

		if reflectValue.CanSet() {
			reflectValue.SetInt(intValue)
			return nil
		}

		return errors.New("default value could not set")
	}
)

func parseInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
