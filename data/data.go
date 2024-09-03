package data

import "time"

type DayHours struct {
	Day       string
	Duration  time.Duration
	Overtime  time.Duration
	Undertime time.Duration
}

func (dh *DayHours) IsOver() bool {
	return dh.Overtime > 0
}

func (dh *DayHours) IsUnder() bool {
	return dh.Undertime > 0
}
