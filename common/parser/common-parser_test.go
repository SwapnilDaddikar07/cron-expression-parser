package parser

import (
	"cron-expression-parser/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommonParser_Parse_Should_ParseCronExpressions(t *testing.T) {
	type scenario struct {
		name          string
		expression    string
		expectedValue []string
	}
	scenarios := []scenario{
		{
			name:          "single minute",
			expression:    "2",
			expectedValue: []string{"2"},
		},
		{
			name:          "list separated values",
			expression:    "2,3,4",
			expectedValue: []string{"2", "3", "4"},
		},
		{
			name:          "range values",
			expression:    "2-5",
			expectedValue: []string{"2", "3", "4", "5"},
		},
		{
			name:          "range values with step values",
			expression:    "2-7/2",
			expectedValue: []string{"2", "4", "6"},
		},
	}

	for _, scenario := range scenarios {
		ctrl := gomock.NewController(t)
		mockParser := mocks.NewMockParser(ctrl)

		commonParser := NewCommonParser()
		actualValue, err := commonParser.Parse(scenario.expression, mockParser)

		assert.Equal(t, scenario.expectedValue, actualValue)
		assert.Nil(t, err)
	}
}

func TestCommonParser_Parse_Should_GetStartAndEndRangeLimitFromTheParserWhen_ExpressionHasStepIdentifierAppliedOn_AnyIdentifier(t *testing.T) {
	expression := "*/2"
	expectedValue := []string{"4", "6", "8", "10"}
	ctrl := gomock.NewController(t)
	mockParser := mocks.NewMockParser(ctrl)

	mockParser.EXPECT().MinAllowedValue().Return(4)
	mockParser.EXPECT().MaxAllowedValue().Return(10)

	commonParser := NewCommonParser()
	actualValue, err := commonParser.Parse(expression, mockParser)

	assert.Equal(t, expectedValue, actualValue)
	assert.Nil(t, err)
}

func TestCommonParser_Parse_Should_GetTheMaxAllowedValueFromTheParser_When_ExpressionHasListSeparatedValues_WithAStepIdentifier(t *testing.T) {
	expression := "2,3,4/2"
	expectedValue := []string{"2", "3", "4", "6", "8", "10"}
	ctrl := gomock.NewController(t)
	mockParser := mocks.NewMockParser(ctrl)

	mockParser.EXPECT().MaxAllowedValue().Return(10)

	commonParser := NewCommonParser()
	actualValue, err := commonParser.Parse(expression, mockParser)

	assert.Equal(t, expectedValue, actualValue)
	assert.Nil(t, err)
}

func TestCommonParser_Parse_Should_GetMinAndMaxAllowedValuesFromTheParser_When_ExpressionHasJustAn_AnyIdentifier(t *testing.T) {
	expression := "*"
	expectedValue := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	ctrl := gomock.NewController(t)
	mockParser := mocks.NewMockParser(ctrl)

	mockParser.EXPECT().MinAllowedValue().Return(1)
	mockParser.EXPECT().MaxAllowedValue().Return(10)

	commonParser := NewCommonParser()
	actualValue, err := commonParser.Parse(expression, mockParser)

	assert.Equal(t, expectedValue, actualValue)
	assert.Nil(t, err)
}

func TestCommonParser_Parse_Should_SetTheStepValueGivenInTheExpression_AndSetTheStartPointAsGivenInThExpression(t *testing.T) {
	expression := "4/2"
	expectedValue := []string{"4", "6"}
	ctrl := gomock.NewController(t)
	mockParser := mocks.NewMockParser(ctrl)
	mockParser.EXPECT().MaxAllowedValue().Return(7)

	commonParser := NewCommonParser()
	actualValue, err := commonParser.Parse(expression, mockParser)

	assert.Equal(t, expectedValue, actualValue)
	assert.Nil(t, err)
}
