package parser

import (
	"cron-expression-parser/common/parser"
	"cron-expression-parser/constant"
	"cron-expression-parser/parser-interface"
	"strings"
)

const (
	firstMonth = 1
	lastMonth  = 12
)

type monthParser struct {
	commonParser parser.CommonParser
}

func NewMonthParser(commonParser parser.CommonParser) parser_interface.Parser {
	return monthParser{commonParser: commonParser}
}

func (m monthParser) Parse(monthComponent string) ([]string, error) {
	result := make([]string, 0)
	for _, individualComponent := range strings.Split(monthComponent, constant.ListValueSeparator) {
		parsedRes, parseErr := m.commonParser.Parse(individualComponent, m)
		if parseErr != nil {
			return nil, parseErr
		}
		result = append(result, parsedRes...)
	}
	return result, nil
}

func (m monthParser) MinAllowedValue() int {
	return firstMonth
}

func (m monthParser) MaxAllowedValue() int {
	return lastMonth
}
