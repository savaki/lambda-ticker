package internal

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	str := "2016-05-04T03:43:44Z"
	when, err := time.Parse(time.RFC3339, str)
	if err != nil {
		t.Errorf("unable to parse time value")
		return
	}

	year, month, day := when.Date()
	if year != 2016 {
		t.Errorf("expected year to be 2016")
	}
	if month != 5 {
		t.Errorf("expected year to be 5")
	}
	if day != 4 {
		t.Errorf("expected year to be 4")
	}

	hour, min, sec := when.Clock()
	if hour != 3 {
		t.Errorf("expected year to be 2016")
	}
	if min != 43 {
		t.Errorf("expected year to be 5")
	}
	if sec != 44 {
		t.Errorf("expected year to be 4")
	}
}
