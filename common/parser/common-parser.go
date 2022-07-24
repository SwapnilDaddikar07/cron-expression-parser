package parser

import (
	"cron-expression-parser/common/representation"
	"cron-expression-parser/constant"
	"cron-expression-parser/parser-interface"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type CommonParser interface {
	Parse(expression string, parser parser_interface.Parser) ([]string, error)
}

type commonParser struct {
}

func NewCommonParser() CommonParser {
	return commonParser{}
}

func (cp commonParser) Parse(expression string, parser parser_interface.Parser) ([]string, error) {

	expressionRepresentation := representation.NewCronExpressionRepresentation()

	if isAllValues(expression) {
		return buildAllValues(expressionRepresentation, parser), nil
	}

	if isStepValueApplicable(expression) {
		modifiedExpression, parseErr := extractIntervalAndModifyExpression(expression, expressionRepresentation)
		if parseErr != nil {
			return nil, parseErr
		}
		expression = modifiedExpression
	}

	if isRangeApplicable(expression) {
		return extractRange(expression, expressionRepresentation)
	}

	if isListValuesApplicable(expression) {
		return extractListValues(expression, expressionRepresentation, parser)
	}

	return processSingularValue(expression, expressionRepresentation, parser)
}

func isAllValues(expression string) bool {
	return expression == constant.AnyValueIdentifier
}

func buildAllValues(expressionRepresentation *representation.ExpressionRepresentation, parser parser_interface.Parser) []string {
	expressionRepresentation.SetStart(parser.MinAllowedValue())
	expressionRepresentation.SetEnd(parser.MaxAllowedValue())
	expressionRepresentation.SetInterval(1)
	return expressionRepresentation.Execute()
}

func isStepValueApplicable(expression string) bool {
	return strings.Contains(expression, constant.StepValuesIdentifier)
}

func extractIntervalAndModifyExpression(expression string, expressionRepresentation *representation.ExpressionRepresentation) (string, error) {
	stepValueIdentifierIndex := strings.LastIndex(expression, constant.StepValuesIdentifier)
	interval, parseErr := strconv.ParseInt(expression[stepValueIdentifierIndex+1:], 10, 0)
	if parseErr != nil {
		errMsg := fmt.Sprintf("cannot parse given expression as it is invalid.Expression - %s", expression)
		return "", errors.New(errMsg)
	}
	expressionRepresentation.SetInterval(int(interval))
	expressionRepresentation.SetIsStepApplicable(true)
	expression = expression[:stepValueIdentifierIndex]
	return expression, nil
}

func isRangeApplicable(expression string) bool {
	return strings.Contains(expression, constant.RangeValuesIdentifier)
}

func extractRange(expression string, expressionRepresentation *representation.ExpressionRepresentation) ([]string, error) {

	rangeValues := strings.Split(expression, constant.RangeValuesIdentifier)
	rangeStart, parseErr := strconv.ParseInt(rangeValues[0], 10, 0)
	rangeEnd, parseErr := strconv.ParseInt(rangeValues[1], 10, 0)

	if parseErr != nil {
		errMsg := fmt.Sprintf("cannot parse given expression as it is invalid. Expression - %s", expression)
		return nil, errors.New(errMsg)
	}

	if !expressionRepresentation.StepValueApplicable() {
		expressionRepresentation.SetInterval(1)
	}
	expressionRepresentation.SetStart(int(rangeStart))
	expressionRepresentation.SetEnd(int(rangeEnd))
	return expressionRepresentation.Execute(), nil
}

func isListValuesApplicable(expression string) bool {
	return strings.Contains(expression, constant.ListValueSeparator)
}

func extractListValues(expression string, expressionRepresentation *representation.ExpressionRepresentation, parser parser_interface.Parser) ([]string, error) {
	listValues := strings.Split(expression, constant.ListValueSeparator)

	for _, val := range listValues {
		_, parseErr := strconv.ParseInt(val, 10, 0)
		if parseErr != nil {
			errMsg := fmt.Sprintf("cannot parse given expression as it is invalid. Expression - %s", expression)
			return nil, errors.New(errMsg)
		}
	}

	if expressionRepresentation.StepValueApplicable() {
		start, _ := strconv.Atoi(listValues[len(listValues)-1])
		listValues = listValues[:len(listValues)-1]
		expressionRepresentation.SetStart(start)
		expressionRepresentation.SetEnd(parser.MaxAllowedValue())
		tmp := expressionRepresentation.Execute()
		listValues = append(listValues, tmp...)
	}
	return listValues, nil
}

func processSingularValue(expression string, expressionRepresentation *representation.ExpressionRepresentation, parser parser_interface.Parser) ([]string, error) {
	if expressionRepresentation.StepValueApplicable() {
		if isAllValues(expression) {
			expressionRepresentation.SetStart(parser.MinAllowedValue())
			expressionRepresentation.SetEnd(parser.MaxAllowedValue())
			return expressionRepresentation.Execute(), nil
		}
		start, parseErr := strconv.ParseInt(expression, 10, 0)
		if parseErr != nil {
			errMsg := fmt.Sprintf("cannot parse given expression as it is invalid. Expression - %s", expression)
			return nil, errors.New(errMsg)
		}
		expressionRepresentation.SetStart(int(start))
		expressionRepresentation.SetEnd(parser.MaxAllowedValue())
		return expressionRepresentation.Execute(), nil
	}
	_, parseErr := strconv.ParseInt(expression, 10, 0)
	if parseErr != nil {
		errMsg := fmt.Sprintf("cannot parse given expression as it is invalid. Expression - %s", expression)
		return nil, errors.New(errMsg)
	}
	return []string{expression}, nil
}
