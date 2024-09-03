package logic

import (
	"fmt"
	"github.com/crossbone-magister/timewlib"
	"hours-day/data"
	"time"
)

var WORKING_HOURS_DAY, _ = time.ParseDuration("8h")

func CalculateDayHours(intervals []timewlib.Interval) ([]data.DayHours, time.Duration, time.Duration, time.Duration) {
	totalPerDay := make(map[string]time.Duration)
	totalHours := time.Duration(0)
	for _, interval := range intervals {
		y, m, d := interval.StartDate()
		key := fmt.Sprintf("%d-%02d-%02d", y, m, d)
		if _, ok := totalPerDay[key]; !ok {
			totalPerDay[key] = 0
		}
		totalPerDay[key] += interval.Duration()
		totalHours += interval.Duration()
	}
	hoursDay := make([]data.DayHours, 0)
	totalOvertime := time.Duration(0)
	totalUndertime := time.Duration(0)
	for day, duration := range totalPerDay {
		hd := data.DayHours{
			Day:       day,
			Duration:  duration,
			Overtime:  time.Duration(0),
			Undertime: time.Duration(0),
		}
		if duration > WORKING_HOURS_DAY {
			hd.Overtime = duration - WORKING_HOURS_DAY
			totalOvertime += hd.Overtime
		}
		if duration < WORKING_HOURS_DAY {
			hd.Undertime = WORKING_HOURS_DAY - duration
			totalUndertime += hd.Undertime
		}
		hoursDay = append(hoursDay, hd)
	}
	return hoursDay, totalHours, totalOvertime, totalUndertime
}
