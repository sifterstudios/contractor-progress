package input

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"

	data "github.com/sifterstudios/gontractor/src/data"
)

func ViewWeek(weeks *map[string]data.Week) {
	if len(*weeks) == 0 {
		fmt.Println("No weeks to view.")
		return
	}

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columntFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Year", "Week", "Normal Hours", "Overtime Hours")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columntFmt)

	for _, week := range *weeks {
		tbl.AddRow(week.Year, week.WeekNumber, week.NormalHours, week.OvertimeHours)
	}

	tbl.Print()
}
