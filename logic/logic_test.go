package logic

import (
	"github.com/crossbone-magister/timewlib"
	"testing"
	"time"
)

func TestCalculateDayHours(t *testing.T) {
	var intervals = []timewlib.Interval{
		*timewlib.NewInterval(9, 0, 10, 0),
		*timewlib.NewInterval(10, 0, 10, 30),
	}
	dayHours, totalTime, _, _ := CalculateDayHours(intervals)
	if len(dayHours) != 1 {
		t.Errorf("dayHours is not collected on the same day: %v", dayHours)
	}
	expected, _ := time.ParseDuration("1h30m")
	if totalTime != expected {
		t.Errorf("totalTime [%s] is not as expected [%s]", totalTime, expected)
	}
}

func TestCalculateDayHoursOvertime(t *testing.T) {
	var intervals = []timewlib.Interval{
		*timewlib.NewInterval(9, 0, 10, 0),
		*timewlib.NewInterval(10, 0, 11, 0),
		*timewlib.NewInterval(11, 0, 12, 0),
		*timewlib.NewInterval(12, 0, 13, 0),
		*timewlib.NewInterval(14, 0, 15, 0),
		*timewlib.NewInterval(15, 0, 16, 0),
		*timewlib.NewInterval(16, 0, 17, 0),
		*timewlib.NewInterval(17, 0, 18, 0),
		*timewlib.NewInterval(18, 0, 18, 30),
	}
	dayHours, totalTime, totalOvertime, _ := CalculateDayHours(intervals)
	if len(dayHours) != 1 {
		t.Errorf("dayHours is not collected on the same day: %v", dayHours)
	}
	expected, _ := time.ParseDuration("8h30m")
	if totalTime != expected {
		t.Errorf("totalTime [%s] is not as expected [%s]", totalTime, expected)
	}
	expected, _ = time.ParseDuration("0h30m")
	if totalOvertime != expected {
		t.Errorf("totalOvertime [%s] is not as expected [%s]", totalOvertime, expected)
	}
}

func TestCalculateDayHoursUndertime(t *testing.T) {
	var intervals = []timewlib.Interval{
		*timewlib.NewInterval(9, 0, 10, 0),
		*timewlib.NewInterval(10, 0, 11, 0),
		*timewlib.NewInterval(11, 0, 12, 0),
		*timewlib.NewInterval(12, 0, 13, 0),
		*timewlib.NewInterval(14, 0, 15, 0),
		*timewlib.NewInterval(15, 0, 16, 0),
		*timewlib.NewInterval(16, 0, 17, 0),
		*timewlib.NewInterval(17, 0, 17, 40),
	}
	dayHours, totalTime, _, totalUndertime := CalculateDayHours(intervals)
	if len(dayHours) != 1 {
		t.Errorf("dayHours is not collected on the same day: %v", dayHours)
	}
	expected, _ := time.ParseDuration("7h40m")
	if totalTime != expected {
		t.Errorf("totalTime [%s] is not as expected [%s]", totalTime, expected)
	}
	expected, _ = time.ParseDuration("20m")
	if totalUndertime != expected {
		t.Errorf("totalUndertime [%s] is not as expected [%s]", totalTime, expected)
	}
}
