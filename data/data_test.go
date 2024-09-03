package data

import (
	"testing"
	"time"
)

func TestIsOver(t *testing.T) {
	dh := DayHours{
		"2024-01-01",
		time.Duration(0),
		time.Duration(10),
		time.Duration(0),
	}
	if !dh.IsOver() {
		t.Errorf("DayHours %v is not overtime", dh)
	}
}

func TestIsUnder(t *testing.T) {
	dh := DayHours{
		"2024-01-01",
		time.Duration(0),
		time.Duration(0),
		time.Duration(10),
	}
	if !dh.IsUnder() {
		t.Errorf("DayHours %v is not undertime", dh)
	}
}

func TestNoOverNoUnder(t *testing.T) {
	dh := DayHours{
		"2024-01-01",
		time.Duration(0),
		time.Duration(0),
		time.Duration(0),
	}
	if dh.IsOver() || dh.IsUnder() {
		t.Errorf("DayHours %v is not over [%t] or under [%t]", dh, dh.IsOver(), dh.IsUnder())
	}
}
