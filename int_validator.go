package gorify

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	betweenElemCountErr      = "between tag should contains floor and ceil values seperated by '-'"
	betweenElemNoIntErr      = "between elements should be int"
	betweenValueIsNotInRange = "given value is not between the range"
)

var (
	between = func(value int, tags reflect.StructTag) error {
		if betweenStr, ok := tags.Lookup("between"); ok {
			minMaxSplitted := strings.Split(betweenStr, "-")
			if len(minMaxSplitted) != 2 {
				return errors.New(betweenElemCountErr)
			}
			min, err := strconv.Atoi(minMaxSplitted[0])
			max, err := strconv.Atoi(minMaxSplitted[1])

			if err != nil {
				return errors.New(betweenElemNoIntErr)
			}

			// value should be converted to int automatically
			// maybe we can send data as byte array too. Because it is simplistic
			// and also for struct types it is more efficient to transfer data in byte format
			// intVal, _ := strconv.Atoi(value)
			if min > value || max < value {
				return errors.New(betweenValueIsNotInRange)
			}
		}
		return nil
	}

	minInt = func(value int, tags reflect.StructTag) error {
		if minIntStr, ok := tags.Lookup("min"); ok {
			min, err := strconv.Atoi(minIntStr)

			if err != nil {
				return errors.New("min tag value should be integer.")
			}

			if min > value {
				return errors.New(fmt.Sprintf("value should be bigger than min: %d", min))
			}
		}
		return nil
	}

	maxInt = func(value int, tags reflect.StructTag) error {
		if maxIntStr, ok := tags.Lookup("max"); ok {
			max, err := strconv.Atoi(maxIntStr)

			if err != nil {
				return errors.New("max tag value should be integer")
			}

			if max < value {
				return errors.New(fmt.Sprintf("value should be lower than max: %d", max))
			}
		}
		return nil
	}
)
