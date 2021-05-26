package gorify

import (
	"reflect"
	"strconv"
	"strings"
)

var (
	between = func(value int, tags reflect.StructTag) error {
		if betweenStr, ok := tags.Lookup("between"); ok {
			minMaxSplitted := strings.Split(betweenStr, "-")
			if len(minMaxSplitted) != 2 {
				panic("min and max should be separated by '-'")
			}
			min, err := strconv.Atoi(minMaxSplitted[0])
			max, err := strconv.Atoi(minMaxSplitted[1])

			if err != nil {
				panic("min and max values should be integer")
			}

			if min > value || max < value {
				return errorMappings["int-between"].(*GorifyErr).objects(value, min, max)
			}
		}
		return nil
	}

	minInt = func(value int, tags reflect.StructTag) error {
		if minIntStr, ok := tags.Lookup("min"); ok {
			min, err := strconv.Atoi(minIntStr)

			if err != nil {
				panic("min value should be integer")
			}

			if min > value {
				return errorMappings["int-min"].(*GorifyErr).objects(value, min)
			}
		}
		return nil
	}

	maxInt = func(value int, tags reflect.StructTag) error {
		if maxIntStr, ok := tags.Lookup("max"); ok {
			max, err := strconv.Atoi(maxIntStr)

			if err != nil {
				panic("max value should be integer")
			}

			if max < value {
				return errorMappings["int-max"].(*GorifyErr).objects(value, max)
			}
		}
		return nil
	}

	setDefaultInt = func(reflectField reflect.StructField, reflectValue reflect.Value) error {
		return setDefaultInt64(reflectField, reflectValue)
	}
)
