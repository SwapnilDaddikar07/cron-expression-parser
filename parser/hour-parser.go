package parser

import (
	"cron-expression-parser/common/parser"
	"cron-expression-parser/constant"
	"cron-expression-parser/parser-interface"
	"strings"
)

const (
	firstHour = 0
	lastHour  = 23
)

type hourParser struct {
	commonParser parser.CommonParser
}

func NewHourParser(commonParser parser.CommonParser) parser_interface.Parser {
	return hourParser{commonParser: commonParser}
}

func (m hourParser) Parse(hourComponent string) ([]string, error) {
	result := make([]string, 0)
	for _, individualComponent := range strings.Split(hourComponent, constant.ListValueSeparator) {
		parsedRes, parseErr := m.commonParser.Parse(individualComponent, m)
		if parseErr != nil {
			return nil, parseErr
		}
		result = append(result, parsedRes...)
	}
	return result, nil
}

func (m hourParser) MinAllowedValue() int {
	return firstHour
}

func (m hourParser) MaxAllowedValue() int {
	return lastHour
}
