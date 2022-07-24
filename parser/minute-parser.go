package parser

import (
	"cron-expression-parser/common/parser"
	"cron-expression-parser/constant"
	"cron-expression-parser/parser-interface"
	"strings"
)

const (
	firstMinute = 0
	lastMinute  = 59
)

type minuteParser struct {
	commonParser parser.CommonParser
}

func NewMinuteParser(commonParser parser.CommonParser) parser_interface.Parser {
	return minuteParser{commonParser: commonParser}
}

func (m minuteParser) Parse(minuteComponent string) ([]string, error) {
	result := make([]string, 0)

	for _, individualComponent := range strings.Split(minuteComponent, constant.ListValueSeparator) {
		parsedRes, parseErr := m.commonParser.Parse(individualComponent, m)
		if parseErr != nil {
			return nil, parseErr
		}
		result = append(result, parsedRes...)
	}
	return result, nil
}

func (m minuteParser) MinAllowedValue() int {
	return firstMinute
}

func (m minuteParser) MaxAllowedValue() int {
	return lastMinute
}
