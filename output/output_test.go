package output

import (
	"hours-day/data"
	"testing"
	"time"
)

func TestFormatTotalDays(t *testing.T) {
	actual := FormatTotalDays([]data.DayHours{
		{"2024-01-01",
			time.Duration(0),
			time.Duration(0),
			time.Duration(0),
		},
	})
	if actual != "Total days: 1" {
		t.Errorf("FormatTotalDays not as expected: %s", actual)
	}
}

func TestFormatTotalOvertime(t *testing.T) {
	duration, _ := time.ParseDuration("1h30m")
	actual := FormatTotalOvertime(duration)
	if actual != "Total overtime: 1h30m0s" {
		t.Errorf("FormatTotalOvertime not as expected: %s", actual)
	}
}

func TestFormatTotalUndertime(t *testing.T) {
	duration, _ := time.ParseDuration("1h30m")
	actual := FormatTotalUndertime(duration)
	if actual != "Total undertime: 1h30m0s" {
		t.Errorf("FormatTotalOvertime not as expected: %s", actual)
	}
}

func TestFormatActualOvertime(t *testing.T) {
	duration, _ := time.ParseDuration("1h30m")
	actual := FormatActualOvertime(duration)
	if actual != "Actual overtime: 1h30m0s" {
		t.Errorf("FormatTotalOvertime not as expected: %s", actual)
	}
}

func TestFormatTotalHours(t *testing.T) {
	duration, _ := time.ParseDuration("1h30m")
	actual := FormatTotalHours(duration)
	if actual != "Total time: 1h30m0s" {
		t.Errorf("FormatTotalHours not as expected: %s", actual)
	}
}

func TestFormatDayHours(t *testing.T) {
	duration, _ := time.ParseDuration("8h")
	durationOver, _ := time.ParseDuration("30m")
	durationUnder, _ := time.ParseDuration("20m")
	actual := FormatDayHours([]data.DayHours{
		{
			"2024-01-01",
			duration,
			time.Duration(0),
			time.Duration(0),
		}, {
			"2024-01-03",
			duration,
			durationOver,
			time.Duration(0),
		}, {
			"2024-01-02",
			duration,
			time.Duration(0),
			durationUnder,
		},
	})
	if actual[0] != "2024-01-01 - 8h0m0s" {
		t.Errorf("Normal day is not as expected: %s", actual[0])
	}
	if actual[1] != "2024-01-02 - 8h0m0s - 20m0s UNDER" {
		t.Errorf("Undertime day is not as expected: %s", actual[1])
	}
	if actual[2] != "2024-01-03 - 8h0m0s - 30m0s OVER" {
		t.Errorf("Overtime day is not as expected: %s", actual[2])
	}

}
