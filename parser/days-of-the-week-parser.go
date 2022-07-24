package parser

import (
	"cron-expression-parser/common/parser"
	"cron-expression-parser/constant"
	"cron-expression-parser/parser-interface"
	"strings"
)

const (
	firstDayOfTheWeek = 0
	lastDayOfTheWeek  = 6
)

type daysOfTheWeekParser struct {
	commonParser parser.CommonParser
}

func NewDaysOfTheWeekParser(commonParser parser.CommonParser) parser_interface.Parser {
	return daysOfTheWeekParser{commonParser: commonParser}
}

func (w daysOfTheWeekParser) IsValid(daysOfTheWeekParser string) bool {
	return true
}

func (w daysOfTheWeekParser) Parse(daysOfTheWeekComponent string) ([]string, error) {
	result := make([]string, 0)
	for _, individualComponent := range strings.Split(daysOfTheWeekComponent, constant.ListValueSeparator) {
		parsedRes, parseErr := w.commonParser.Parse(individualComponent, w)
		if parseErr != nil {
			return nil, parseErr
		}
		result = append(result, parsedRes...)
	}
	return result, nil
}

func (w daysOfTheWeekParser) MinAllowedValue() int {
	return firstDayOfTheWeek
}

func (w daysOfTheWeekParser) MaxAllowedValue() int {
	return lastDayOfTheWeek
}
