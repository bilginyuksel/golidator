package gorify

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	maxInt64 = func(value int64, tags reflect.StructTag) error {
		if maxStr, ok := tags.Lookup("max"); ok {
			maxInt64, err := parseInt64(maxStr)
			if err != nil {
				return errors.New("please write max constraint as int64")
			}
			if maxInt64 < value {
				return errors.New(fmt.Sprintf("the value should be smaller than %d", maxInt64))
			}
		}
		return nil
	}

	minInt64 = func(value int64, tags reflect.StructTag) error {
		if minStr, ok := tags.Lookup("min"); ok {
			minInt64, err := parseInt64(minStr)
			if err != nil {
				return errors.New("please write min constraint as int64")
			}
			if minInt64 > value {
				return errors.New(fmt.Sprintf("the value should be bigger than %d", minInt64))
			}
		}
		return nil
	}

	betweenInt64 = func(value int64, tags reflect.StructTag) error {
		if betweenStr, ok := tags.Lookup("between"); ok {
			betweenStrSplitted := strings.Split(betweenStr, ",")

			var err error
			minInt64, err := parseInt64(betweenStrSplitted[0])
			maxInt64, err := parseInt64(betweenStrSplitted[1])
			if err != nil {
				return errors.New("please write int64 separated with ',' inside between tag")
			}
			if minInt64 > value || maxInt64 < value {
				return errors.New(fmt.Sprintf("value should be between (%d, %d)", minInt64, maxInt64))
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
