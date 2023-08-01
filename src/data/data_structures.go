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

type Stats struct {
	TotalWeeks            int
	TotalHours            float64
	TotalOvertimeHours    float64
	TotalVacationDays     int
	TotalNationalHolidays int
	TotalChildcareDays    int
	TotalSickDays         int
	PctCompleted          float64
	DaysLeft              int
	HoursLeft             float64
}
