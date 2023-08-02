package input

import (
	"fmt"

	data "github.com/sifterstudios/gontractor/src/data"
)

func ChangeWeek(weeks *map[string]data.Week) {
	fmt.Println("Changing a week...")
	week := data.Week{}
	week = (*weeks)[fmt.Sprintf("%d-%d", week.Year, week.WeekNumber)]
	fmt.Println("Are there any special things needing to be accounted for this week? (y/n)")
	fmt.Scan(&input)
	if input == "y" {
		getSpecialThings(week, weeks)
	}
	// Add week to weeks
	(*weeks)[fmt.Sprintf("%d-%d", week.Year, week.WeekNumber)] = week
}
