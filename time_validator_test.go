package gorify

import (
	"testing"
	"time"
)

type TestTime struct {
	T time.Time `default:"now,add-d9,utc,round"`
}

func TestTimeDefault_now(t *testing.T) {
	tt := &TestTime{}
	Validate(tt)
	t.Logf("%+v", tt)
}
