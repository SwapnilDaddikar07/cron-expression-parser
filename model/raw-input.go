package model

type RawInput struct {
	Minute        string
	Hour          string
	DayOfTheMonth string
	Month         string
	DayOfTheWeek  string
	Command       string
}

func NewRawInput(minute, hour, dayOfTheMonth, month, dayOfTheWeek, command string) RawInput {
	return RawInput{Minute: minute, Hour: hour, DayOfTheMonth: dayOfTheMonth, Month: month, DayOfTheWeek: dayOfTheWeek, Command: command}
}
