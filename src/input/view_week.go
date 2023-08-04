package input

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"

	data "github.com/sifterstudios/gontractor/src/data"
)

var (
	yearNumber int
	weekNumber int
)

func ViewWeek(weeks *map[string]data.Week) {
	if len(*weeks) == 0 {
		fmt.Println("No weeks to view.")
		return
	}
	inputString := ""

	fmt.Println("Using this year. Press n to specify, or any other key to continue.")
	fmt.Scan(&inputString)
	if inputString == "n" {
		fmt.Println("Which year would you like to view?")
		fmt.Scan(&yearNumber)
	}

	fmt.Println("Which week would you like to view?")
	fmt.Scan(&weekNumber)

	fmt.Printf("Using year %d, week %d\n", yearNumber, weekNumber)

	// Get Week using key combinging year and week number
	week, found := (*weeks)[fmt.Sprintf("%d-%d", yearNumber, weekNumber)]
	if !found {
		fmt.Println("Week not found.")
		return
	}

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columntFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Year", "Week", "Normal Hours", "Overtime Hours", "Vacation Days", "Sick Days", "National Holidays", "Childcare Days")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columntFmt)

	tbl.AddRow(week.Year, week.WeekNumber, week.NormalHours, week.OvertimeHours, week.VacationDays, week.SickDays, week.NationalHolidays, week.ChildcareDays)

	tbl.Print()
}
