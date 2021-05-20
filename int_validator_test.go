package gorify

import "testing"

type IntStruct struct {
	SampleInt int `min:"100" max:"5000"`
	Age       int `json:"age" between:"20-100"`
}

func TestValidation_IntField(t *testing.T) {
	testCases := []struct {
		scenario string
		is       *IntStruct
		expected bool
	}{
		{
			scenario: "equal to min range",
			is:       &IntStruct{Age: 20, SampleInt: 1000},
			expected: true,
		},
		{
			scenario: "between min-max range",
			is:       &IntStruct{Age: 55, SampleInt: 120},
			expected: true,
		},
		{
			scenario: "equal to max range",
			is:       &IntStruct{Age: 100, SampleInt: 4830},
			expected: true,
		},
		{
			scenario: "smaller than min range",
			is:       &IntStruct{Age: 10, SampleInt: 500},
			expected: false,
		},
		{
			scenario: "greater than max range",
			is:       &IntStruct{Age: 120, SampleInt: 500},
			expected: false,
		},
		{
			scenario: "smaller than min constraint",
			is:       &IntStruct{Age: 50, SampleInt: 50},
			expected: false,
		},
		{
			scenario: "bigger than max constraint",
			is:       &IntStruct{Age: 50, SampleInt: 5005},
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.scenario, func(t *testing.T) {
			if err := Validate(tC.is); err != nil && tC.expected {
				t.Errorf("failed, err: %v", err)
			}
		})
	}
}
