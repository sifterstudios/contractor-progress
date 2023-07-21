package data

type Week struct {
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
