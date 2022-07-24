package parser

import (
	"cron-expression-parser/model"
	"cron-expression-parser/parser-interface"
	"errors"
	"strings"
)

const (
	separator                            string = " "
	allowedCronExpressionComponentsCount int    = 6
)

type CronExpressionParser interface {
	Parse(expression string) (model.ParsedCronExpression, error)
	BuildRawCronExpressionComponents(expression string) (model.RawInput, error)
}

type cronExpressionParser struct {
	minuteParser         parser_interface.Parser
	hourParser           parser_interface.Parser
	daysOfTheMonthParser parser_interface.Parser
	monthParser          parser_interface.Parser
	daysOfTheWeekParser  parser_interface.Parser
}

func NewCronExpressionParser(minuteParser parser_interface.Parser,
	hourParser parser_interface.Parser,
	daysOfTheMonthParser parser_interface.Parser,
	monthParser parser_interface.Parser,
	daysOfTheWeekParser parser_interface.Parser) CronExpressionParser {
	return cronExpressionParser{minuteParser: minuteParser,
		hourParser:           hourParser,
		daysOfTheMonthParser: daysOfTheMonthParser,
		monthParser:          monthParser,
		daysOfTheWeekParser:  daysOfTheWeekParser}
}

func (c cronExpressionParser) Parse(expression string) (model.ParsedCronExpression, error) {
	rawCronExpressionComponents, err := c.BuildRawCronExpressionComponents(expression)
	if err != nil {
		return model.ParsedCronExpression{}, err
	}
	return c.parse(rawCronExpressionComponents)
}

func (c cronExpressionParser) BuildRawCronExpressionComponents(expression string) (model.RawInput, error) {
	individualComponents := strings.Split(expression, separator)
	if len(individualComponents) != allowedCronExpressionComponentsCount {
		return model.RawInput{}, errors.New("invalid cron expression")
	}
	return model.NewRawInput(individualComponents[0],
		individualComponents[1],
		individualComponents[2],
		individualComponents[3],
		individualComponents[4],
		individualComponents[5]), nil
}

func (c cronExpressionParser) parse(rawInput model.RawInput) (parsedExpression model.ParsedCronExpression, err error) {
	var parsedCronExpression model.ParsedCronExpression

	parsedMinutes, err := c.minuteParser.Parse(rawInput.Minute)
	if err != nil {
		return parsedExpression, err
	}
	parsedCronExpression.Minutes = parsedMinutes

	parsedHours, err := c.hourParser.Parse(rawInput.Hour)
	if err != nil {
		return parsedExpression, err
	}
	parsedCronExpression.Hours = parsedHours

	parsedDaysOfTheMonth, err := c.daysOfTheMonthParser.Parse(rawInput.DayOfTheMonth)
	if err != nil {
		return parsedExpression, err
	}
	parsedCronExpression.DaysOfTheMonth = parsedDaysOfTheMonth

	parsedMonths, err := c.monthParser.Parse(rawInput.Month)
	if err != nil {
		return parsedExpression, err
	}
	parsedCronExpression.Months = parsedMonths

	parsedDaysOfTheWeek, err := c.daysOfTheWeekParser.Parse(rawInput.DayOfTheWeek)
	if err != nil {
		return parsedExpression, err
	}
	parsedCronExpression.DaysOfTheWeek = parsedDaysOfTheWeek

	parsedCronExpression.Command = rawInput.Command
	return parsedCronExpression, nil

}
