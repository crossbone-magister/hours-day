package main

import (
	"fmt"
	"github.com/crossbone-magister/timewlib"
	"hours-day/logic"
	"hours-day/output"
	"os"
)

func main() {
	parsed, err := timewlib.Parse(os.Stdin)
	if err == nil {
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
			printErrorAndExit(err)
		}
	} else {
		printErrorAndExit(err)
	}
}

func printErrorAndExit(err error) {
	fmt.Printf("Error while reading timewarrior input: %s\n", err)
	os.Exit(1)
}
