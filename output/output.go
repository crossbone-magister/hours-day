package output

import (
	"fmt"
	"hours-day/data"
	"sort"
	"time"
)

func FormatDayHours(dayHours []data.DayHours) []string {
	output := make([]string, 0)
	sort.Slice(dayHours, func(i, j int) bool {
		return dayHours[i].Day < dayHours[j].Day
	})
	for _, day := range dayHours {
		formattedDay := fmt.Sprintf("%s - %s", day.Day, day.Duration)
		if day.IsOver() {
			formattedDay = fmt.Sprintf("%s - %s OVER", formattedDay, day.Overtime)
		}
		if day.IsUnder() {
			formattedDay = fmt.Sprintf("%s - %s UNDER", formattedDay, day.Undertime)
		}
		output = append(output, formattedDay)
	}
	return output
}

func FormatTotalHours(hours time.Duration) string {
	return formatTotal("time", hours)
}

func FormatTotalOvertime(hours time.Duration) string {
	return formatTotal("overtime", hours)
}

func FormatTotalUndertime(hours time.Duration) string {
	return formatTotal("undertime", hours)
}

func FormatTotalDays(days []data.DayHours) string {
	return fmt.Sprintf("Total days: %d", len(days))
}

func FormatActualOvertime(hours time.Duration) string {
	return fmt.Sprintf("Actual overtime: %s", hours)
}

func formatTotal(name string, hours time.Duration) string {
	return fmt.Sprintf("Total %s: %s", name, hours)
}
