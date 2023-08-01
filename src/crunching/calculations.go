package crunching

import data "github.com/sifterstudios/gontractor/src/data"

func getTotalHours(weeks map[string]data.Week) float64 {
	totalHours := 0.0
	for _, week := range weeks {
		totalHours += week.NormalHours + week.OvertimeHours
	}
	return totalHours
}
