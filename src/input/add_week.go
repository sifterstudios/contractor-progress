package input

import (
	"fmt"

	data "github.com/sifterstudios/gontractor/src/data"
)

var (
	inputString string
	err         error
)

func AddWeek(weeks *map[string]data.Week) {
	if weeks == nil {
		*weeks = make(map[string]data.Week)
	}

	week := data.Week{}

	fmt.Println("Adding a week...")
	currentYear, currentWeek := data.GetCurrentYearAndWeek()
	fmt.Printf("Using year %d, and week #%d. Press 'y' to confirm, 'n' to specify. ", currentYear, currentWeek)
	fmt.Scan(&inputString)

	if inputString == "n" {
		getCustomYearAndWeek(&week)
	} else {
		week.Year = currentYear
		week.WeekNumber = currentWeek
	}
	fmt.Println("Are there any special things needing to be accounted for this week? (y/n) ")
	fmt.Scan(&inputString)

	week.NormalHours, err = PromptFloat("How many regular hours did you work this week? ")
	if err != nil {
		AddWeek(weeks)
	}
	week.OvertimeHours, err = PromptFloat("How many overtime hours did you work this week? ")
	if err != nil {
		AddWeek(weeks)
	}
	if inputString == "y" {
		getSpecialThings(&week, &weeks)
	}

	summarizeWeek(&week)
	fmt.Println("Press any key to confirm, 'n' to re-enter. ")
	if inputString == "n" {
		AddWeek(weeks)
	}

	(*weeks)[fmt.Sprintf("%d-%d", week.Year, week.WeekNumber)] = week
}

func summarizeWeek(week *data.Week) {
	fmt.Printf("Week %d-%d\n", week.Year, week.WeekNumber)
	fmt.Printf("Regular hours: %f\n", week.NormalHours)
	fmt.Printf("Overtime hours: %f\n", week.OvertimeHours)
	fmt.Printf("Vacation days: %d\n", week.VacationDays)
	fmt.Printf("National holidays: %d\n", week.NationalHolidays)
	fmt.Printf("Childcare days: %d\n", week.ChildcareDays)
	fmt.Printf("Sick days: %d\n", week.SickDays)
}

func getSpecialThings(week *data.Week, weeks **map[string]data.Week) {
	week.VacationDays, err = PromptInt("How many vacation days did you take this week? ")
	if err != nil {
		AddWeek(*weeks)
	}
	week.NationalHolidays, err = PromptInt("How many national holidays did you take this week? ")
	if err != nil {
		AddWeek(*weeks)
	}
	week.ChildcareDays, err = PromptInt("How many childcare days did you take this week? ")
	if err != nil {
		AddWeek(*weeks)
	}
	week.SickDays, err = PromptInt("How many sick days did you take this week? ")
	if err != nil {
		AddWeek(*weeks)
	}
}

func getCustomYearAndWeek(week *data.Week) {
	fmt.Print("Enter year: ")
	fmt.Scan(&week.Year)
	fmt.Print("Enter week: ")
	fmt.Scan(&week.WeekNumber)
}
