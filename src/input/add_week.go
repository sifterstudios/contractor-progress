package input

import (
	"fmt"
	"time"

	data "github.com/sifterstudios/gontractor/src/data"
)

var (
	input string
	err   error
)

func AddWeek(weeks *map[string]data.Week) {
	week := data.Week{}

	fmt.Println("Adding a week...")
	currentYear, currentWeek := getCurrentWeek()
	fmt.Printf("Using year %d, and week #%d. Press return to confirm, any other key to specify.", currentYear, currentWeek)
	fmt.Scan(&input)
	if input != " " {
		getCustomYearAndWeek(&week)
	} else {
		week.Year = currentYear
		week.WeekNumber = currentWeek
	}
	fmt.Println("Are there any special things needing to be accounted for this week? (y/n)")
	fmt.Scan(&input)

	week.NormalHours, err = PromptFloat("How many regular hours did you work this week?")
	if err != nil {
		AddWeek(weeks)
	}
	week.OvertimeHours, err = PromptFloat("How many overtime hours did you work this week?")
	if err != nil {
		AddWeek(weeks)
	}
	if input == "y" {
		getSpecialThings(week, weeks)
	}

	// Add week to weeks
	(*weeks)[fmt.Sprintf("%d-%d", week.Year, week.WeekNumber)] = week
}

func getSpecialThings(week data.Week, weeks *map[string]data.Week) {
	week.VacationDays, err = PromptInt("How many vacation days did you take this week?")
	if err != nil {
		AddWeek(weeks)
	}
	week.NationalHolidays, err = PromptInt("How many national holidays did you take this week?")
	if err != nil {
		AddWeek(weeks)
	}
	week.ChildcareDays, err = PromptInt("How many childcare days did you take this week?")
	if err != nil {
		AddWeek(weeks)
	}
	week.SickDays, err = PromptInt("How many sick days did you take this week?")
	if err != nil {
		AddWeek(weeks)
	}
}

func getCustomYearAndWeek(week *data.Week) {
	fmt.Print("Enter year: ")
	fmt.Scan(&week.Year)
	fmt.Print("Enter week: ")
	fmt.Scan(&week.WeekNumber)
}

func getCurrentWeek() (int, int) {
	return time.Now().ISOWeek()
}
