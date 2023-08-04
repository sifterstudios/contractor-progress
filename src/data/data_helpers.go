package data

import "time"

func GetCurrentYearAndWeek() (int, int) {
	return time.Now().ISOWeek()
}
