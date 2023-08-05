package crunching

import (
	"math"
	"time"

	data "github.com/sifterstudios/gontractor/src/data"
)

func GetTotalHours(weeks *map[string]data.Week) float64 {
	totalHours := 0.0
	for _, week := range *weeks {
		totalHours += week.NormalHours + week.OvertimeHours
	}
	return totalHours
}

func GetPercentComplete(weeks *map[string]data.Week, goal *data.Goal) float64 {
	totalHours := GetTotalHours(weeks)
	return totalHours / goal.TotalContractHours
}

func GetTimeLeft(weeks *map[string]data.Week, goal *data.Goal) (int, float64) {
	totalHours := GetTotalHours(weeks)
	hoursLeft := goal.TotalContractHours - totalHours
	daysLeft := math.Floor(hoursLeft / 7.5)
	hoursLeft = hoursLeft - (daysLeft * 7.5)
	return int(daysLeft), hoursLeft
}

func GetTotalSickDays(weeks *map[string]data.Week) int {
	totalSickDays := 0
	for _, week := range *weeks {
		totalSickDays += week.SickDays
	}
	return totalSickDays
}

func GetTotalVacationDays(weeks *map[string]data.Week) int {
	totalVacationDays := 0
	for _, week := range *weeks {
		totalVacationDays += week.VacationDays
	}
	return totalVacationDays
}

func GetTotalChildcareDays(weeks *map[string]data.Week) int {
	totalChildCareDays := 0
	for _, week := range *weeks {
		totalChildCareDays += week.ChildcareDays
	}
	return totalChildCareDays
}

func GetTotalNationalHolidays(weeks *map[string]data.Week) int {
	nationalHolidays := 0
	for _, week := range *weeks {
		nationalHolidays += week.NationalHolidays
	}
	return nationalHolidays
}

func GetEstimatedCompletionDate(weeks *map[string]data.Week, goal *data.Goal) time.Time {
	totalHours := GetTotalHours(weeks)
	estimatedHoursLeft := goal.TotalContractHours - totalHours
	estimatedDaysLeft := math.Floor(estimatedHoursLeft / 7.5)
	estimatedHoursLeft = estimatedHoursLeft - (estimatedDaysLeft * 7.5)
	estimatedCompletionDate := time.Now()

	for estimatedDaysLeft > 0 {
		if estimatedCompletionDate.Weekday() == 0 || estimatedCompletionDate.Weekday() == 6 {
			estimatedCompletionDate = estimatedCompletionDate.AddDate(0, 0, 1)
		} else {
			estimatedCompletionDate = estimatedCompletionDate.AddDate(0, 0, 1)
			estimatedDaysLeft--
		}
	}

	return estimatedCompletionDate
}
