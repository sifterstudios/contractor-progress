package gontractor

func getTotalHours(weeks map[string]Week) float64 {
	totalHours := 0.0
	for _, week := range weeks {
		totalHours += week.NormalHours + week.OvertimeHours
	}
	return totalHours
}
