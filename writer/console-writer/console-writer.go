package console_writer

import (
	model2 "cron-expression-parser/model"
	"fmt"
	"strings"
)

const (
	leftJustifiedPadding = "%-14s"
	separator            = " "
)

type ConsoleWriter struct {
}

func NewConsoleWriter() ConsoleWriter {
	return ConsoleWriter{}
}

func (c *ConsoleWriter) Write(expression model2.ParsedCronExpression) error {
	formatter := leftJustifiedPadding + " %s \n"
	fmt.Printf(formatter, "minute", strings.Join(expression.Minutes, separator))
	fmt.Printf(formatter, "hour", strings.Join(expression.Hours, separator))
	fmt.Printf(formatter, "day of month", strings.Join(expression.DaysOfTheMonth, separator))
	fmt.Printf(formatter, "month ", strings.Join(expression.Months, separator))
	fmt.Printf(formatter, "day of week ", strings.Join(expression.DaysOfTheWeek, separator))
	fmt.Println()
	return nil
}
