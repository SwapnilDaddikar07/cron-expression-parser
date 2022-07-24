package parser

import (
	"cron-expression-parser/common/parser"
	"cron-expression-parser/constant"
	"cron-expression-parser/parser-interface"
	"strings"
)

const (
	firstDayOfTheMonth = 1
	lastDayOfTheMonth  = 31
)

type dayOfTheMonthParser struct {
	commonParser parser.CommonParser
}

func NewDayOfTheMonthParser(commonParser parser.CommonParser) parser_interface.Parser {
	return dayOfTheMonthParser{commonParser: commonParser}
}

func (m dayOfTheMonthParser) Parse(daysOfTheMonthComponent string) ([]string, error) {
	result := make([]string, 0)
	for _, individualComponent := range strings.Split(daysOfTheMonthComponent, constant.ListValueSeparator) {
		parsedRes, parseErr := m.commonParser.Parse(individualComponent, m)
		if parseErr != nil {
			return nil, parseErr
		}
		result = append(result, parsedRes...)
	}
	return result, nil
}

func (m dayOfTheMonthParser) MinAllowedValue() int {
	return firstDayOfTheMonth
}

func (m dayOfTheMonthParser) MaxAllowedValue() int {
	return lastDayOfTheMonth
}
