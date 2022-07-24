package model

type RawCronExpressionComponents struct {
	Minute        string
	Hour          string
	DayOfTheMonth string
	Month         string
	DayOfTheWeek  string
}

func NewRawCronExpressionComponent(minute, hour, dayOfTheMonth, month, dayOfTheWeek string) RawCronExpressionComponents {
	return RawCronExpressionComponents{Minute: minute, Hour: hour, DayOfTheMonth: dayOfTheMonth, Month: month, DayOfTheWeek: dayOfTheWeek}
}
