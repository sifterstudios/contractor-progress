package data

type Week struct {
	WeekNumber       string
	Year             int
	NormalHours      float64
	OvertimeHours    float64
	VacationDays     int
	NationalHolidays int
	ChildcareDays    int
	SickDays         int
}

type Goal struct {
	TotalContractHours float64
}

type FileData struct {
	Weeks map[string]Week
	Goal  Goal
}
