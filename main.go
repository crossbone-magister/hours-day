package main

import (
	"fmt"
	"hours-day/logic"
	"hours-day/output"
	"os"

	"github.com/crossbone-magister/timewlib"
)

func main() {
	parsed, err := timewlib.Parse(os.Stdin)
	if err == nil {
		timewlib.ExitOnNoData(parsed.Intervals, parsed.Configuration)
		timewlib.SetupLogging(parsed.Configuration)
		intervals, err := timewlib.Process(parsed.Intervals)
		hoursDay, totalHours, totalOvertime, totalUndertime := logic.CalculateDayHours(intervals)
		if err == nil {
			for _, row := range output.FormatDayHours(hoursDay) {
				fmt.Println(row)
			}
			fmt.Println(output.FormatTotalDays(hoursDay))
			fmt.Println(output.FormatTotalHours(totalHours))
			fmt.Println(output.FormatTotalOvertime(totalOvertime))
			fmt.Println(output.FormatTotalUndertime(totalUndertime))
			fmt.Println(output.FormatActualOvertime(totalOvertime - totalUndertime))
		} else {
			timewlib.ExitIfError(err)
		}
	} else {
		timewlib.ExitIfError(err)
	}
}
