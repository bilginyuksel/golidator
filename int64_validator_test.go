package gorify

import "testing"

type Int64Struct struct {
	FileSizeInBytes int64 `max:"24654118912" min:"367001600"` // min 350MB, max:23.5GB
	Timestamp       int64 `min:"55100" max:"128345"`
}

type Int64SetDefault struct {
	FileSizeLimitBytes int64 `default:"24654118912"`
}

func TestInt64Validation(t *testing.T) {
	testCases := []struct {
		desc     string
		is       *Int64Struct
		expected bool
	}{
		{
			desc: "Timestamp smaller than min constraint",
			is:   &Int64Struct{FileSizeInBytes: 1024 * 1024 * 400, Timestamp: 38212},
		},
		{
			desc: "Timestamp greater than max constraint",
			is:   &Int64Struct{FileSizeInBytes: 1024 * 1024 * 400, Timestamp: 158321},
		},
		{
			desc: "FileSizeInBytes smaller than min constraint 300MB",
			is:   &Int64Struct{FileSizeInBytes: 1024 * 1024 * 300, Timestamp: 111555},
		},
		{
			desc: "FileSizeInBytes greater than min constraint 24GB",
			is:   &Int64Struct{FileSizeInBytes: 1024 * 1024 * 24000, Timestamp: 111555},
		},
		{
			desc:     "All in range",
			is:       &Int64Struct{FileSizeInBytes: 1024 * 1024 * 3000, Timestamp: 111555},
			expected: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if err := Validate(tC.is); err != nil && tC.expected {
				t.Errorf("failed, err: %v", err)
			} else {
				t.Log(err)
			}
		})
	}
}

func TestInt64Validation_SetDefaultValue(t *testing.T) {
	testObj := &Int64SetDefault{}
	t.Logf("test obj before validation: %v", testObj)
	if err := Validate(testObj); err != nil {
		t.Errorf("validation gives an error, but it shouldn't")
	}
	t.Logf("test obj after validation: %v", testObj)
	if testObj.FileSizeLimitBytes != 24654118912 {
		t.Errorf("default value could not set successfully")
	}
}
