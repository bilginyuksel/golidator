package gorify

import (
	"errors"
	"log"
	"reflect"
	"time"
)

var (
	// setDefaultTime you can use default tag with the time struct to set a default value to time value.
	// providing an expression to default value you can process relatively complex time creation process.
	// Example: default:"now,add-10d,utc,round"
	// The expression above will execute the statements below
	// 	t := time.Now()
	//	t.Add(time.Day * 10)
	//	t.UTC()
	//	t.Round()
	setDefaultTime = func(reflectField reflect.StructField, reflectValue reflect.Value) error {
		tags := reflectField.Tag

		expression, exists := tags.Lookup("default")
		if !exists {
			return nil
		}

		if reflectValue.CanSet() {
			// strategy := parseTimeExpression(expression)
			// customTime := strategy.build()
			customTime := time.Now()
			log.Printf("expression: %v", expression)

			reflectCustom := reflect.ValueOf(customTime)

			reflectValue.Set(reflectCustom)
			return nil
		}

		return errors.New("could not set default value to time")
	}
)

type timeBuildStrategy struct {
}

func parseTimeExpression(expression string) timeBuildStrategy {
	return timeBuildStrategy{}
}

func (ts timeBuildStrategy) build() time.Time {
	return time.Time{}
}
