package model

type ParsedCronExpression struct {
	Minutes        []string
	Hours          []string
	DaysOfTheMonth []string
	Months         []string
	DaysOfTheWeek  []string
	Command        string
}
