package gorify

import (
	"errors"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	// NOTE: check int overflow for add method

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
			log.Printf("expression: %v", expression)
			strategy := parseTimeExpression(expression)
			log.Printf("strategy: %v", strategy)
			customTime := strategy.build()

			reflectCustom := reflect.ValueOf(customTime)

			reflectValue.Set(reflectCustom)
			return nil
		}

		return errors.New("could not set default value to time")
	}
)

var (
	validInitStatements = map[string]bool{
		"now":   true,
		"empty": true,
		"-":     true,
		"":      true,
	}

	// addDate should be ordered as 1 because of the conflict with add
	// if we check with the help of contains method then this order will be safe
	validEditStatements = []string{"addDate", "sub", "add"}

	// Create the same logic with the edit
	validUtilityStatements = []string{"utc", "local", "round"}
)

type editTime interface {
	method() string

	// year, month, day
	date() (int, int, int)

	// base duration
	duration() time.Duration
}

type timeAddSub struct {
	methodName    string
	durationAlias string
}

type timeAddDate struct {
	methodName string
	day        int
	month      int
	year       int
}

// timeAddDate
func (t timeAddDate) date() (int, int, int) {
	return t.year, t.month, t.day
}
func (t timeAddDate) duration() time.Duration {
	return 0
}
func (t timeAddDate) method() string {
	return t.methodName
}

// timeAddSub
func (t timeAddSub) date() (int, int, int) {
	return 0, 0, 0
}
func (t timeAddSub) duration() time.Duration {
	dtion, err := time.ParseDuration(t.durationAlias)

	if err != nil {
		log.Printf("wrong duration provided, duration should be (300ms, -1.5h or 2h45m. Valid time units are ns, us (or µs), ms, s, m, h)")
		return 0
	}

	return dtion
}
func (t timeAddSub) method() string {
	return t.methodName
}

type timeBuildStrategy struct {
	init      string
	edits     []editTime
	utilities []string
}

func parseTimeExpression(expression string) timeBuildStrategy {
	var (
		initStatement  = ""
		editStatements = []editTime{}
		utilites       = []string{}
	)

	tokens := strings.Split(expression, ",")
	log.Printf("tokens: %v", tokens)
	// now,empty-init,-
	// add-(d5,h2,y1...), sub-(d5,h2,y1...), addDate-2030y15m10d
	// utc
	if len(tokens) > 1 && validInitStatements[tokens[0]] {
		initStatement = tokens[0]
	}

	for i := 1; i < len(tokens); i++ {
		log.Printf("tokens[%d]: %v", i, tokens[i])
		// check utility and edit strategies here
		if contains(validEditStatements, tokens[i]) {
			// parse edit strategy
			editStt := strings.Split(tokens[i], "-")
			log.Printf("editStt: %v", editStt)
			if len(editStt) != 2 {
				log.Printf("wrong argument, method name and param should be provided and seperated by '-', edit statement skipped")
				continue
			}

			method := editStt[0]
			param := editStt[1]

			if method == "addDate" {
				dateParam := parseTimeAddDate(param)
				editStatements = append(editStatements, dateParam)
			} else {
				editStatements = append(editStatements, timeAddSub{methodName: method, durationAlias: param})
			}

		} else if contains(validUtilityStatements, tokens[i]) {
			utilites = append(utilites, tokens[i])
		}
	}

	return timeBuildStrategy{
		init:      initStatement,
		edits:     editStatements,
		utilities: utilites,
	}
}
func contains(stts []string, str string) bool {
	for _, stt := range stts {
		if strings.Contains(str, stt) {
			return true
		}
	}
	return false
}

func parseTimeAddDate(str string) timeAddDate {
	var (
		curr  = "0"
		day   = 0
		month = 0
		year  = 0
	)

	for _, char := range str {
		currint, err := strconv.Atoi(curr)
		if err != nil {
			log.Printf("illegal parameter")
			return timeAddDate{}
		}

		if char == 'd' {
			day = currint
		} else if char == 'm' {
			month = currint
		} else if char == 'y' {
			year = currint
		}

		curr += string(char)
	}

	return timeAddDate{
		methodName: "addDate",
		day:        day,
		month:      month,
		year:       year,
	}
}

func (ts timeBuildStrategy) build() time.Time {
	var customTime time.Time

	if ts.init == "now" {
		customTime = time.Now()
	} else {
		customTime = time.Time{}
	}

	for _, edit := range ts.edits {
		if edit.method() == "addDate" {
			customTime = customTime.AddDate(edit.date())
		} else if edit.method() == "add" {
			customTime = customTime.Add(edit.duration())
		}
		// else if edit.method() == "sub" {
		// 	customTime = customTime.Sub(edit.duration())
		// }
	}

	return customTime
}
