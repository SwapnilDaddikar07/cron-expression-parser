package parser

import (
	"cron-expression-parser/mocks"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDaysOfTheMonthParser_Parse_ShouldCallTheCommonParser_Once_AndReturnCombinedValues(t *testing.T) {
	type scenario struct {
		name                 string
		daysOfTheMonthParser string
		expectedValue        []string
	}
	scenarios := []scenario{
		{
			name:                 "single daysOfTheMonth",
			daysOfTheMonthParser: "5",
			expectedValue:        []string{"5"},
		},
		{
			name:                 "range of daysOfTheMonth",
			daysOfTheMonthParser: "1-2",
			expectedValue:        []string{"1", "2"},
		},
		{
			name:                 "range with step",
			daysOfTheMonthParser: "1-3/1",
			expectedValue:        []string{"1", "2", "3"},
		},
	}

	for _, scenario := range scenarios {
		ctrl := gomock.NewController(t)
		mockCommonParser := mocks.NewMockCommonParser(ctrl)
		mockCommonParser.EXPECT().Parse(scenario.daysOfTheMonthParser, gomock.Any()).Return(scenario.expectedValue, nil)

		dayOfTheMonthParser := NewDayOfTheMonthParser(mockCommonParser)
		actualValue, err := dayOfTheMonthParser.Parse(scenario.daysOfTheMonthParser)

		assert.Nil(t, err)
		assert.Equal(t, scenario.expectedValue, actualValue)
	}
}

func TestTestDaysOfTheMonthParser_Parse_ShouldCallTheCommonParserThrice_AsThereAreThreeListValues(t *testing.T) {
	daysOfTheMonth := "3-4,6,7"
	ctrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(ctrl)
	mockCommonParser.EXPECT().Parse("3-4", gomock.Any()).Return([]string{"3", "4"},nil)
	mockCommonParser.EXPECT().Parse("6", gomock.Any()).Return([]string{"6"},nil)
	mockCommonParser.EXPECT().Parse("7", gomock.Any()).Return([]string{"7"},nil)

	daysOfTheMonthParser := NewDayOfTheMonthParser(mockCommonParser)
	actualValue, err := daysOfTheMonthParser.Parse(daysOfTheMonth)

	expectedValue := []string{"3", "4", "6", "7"}
	assert.Equal(t, expectedValue, actualValue)
	assert.Nil(t, err)
}

func TestTestDaysOfTheMonthParser_MaxAllowedValue_ShouldReturn_31(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(mockCtrl)

	daysOfTheMonthParser := NewDayOfTheMonthParser(mockCommonParser)
	assert.Equal(t, 31, daysOfTheMonthParser.MaxAllowedValue())
}

func TestTestDaysOfTheMonthParser_MinAllowedValue_ShouldReturn_1(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(mockCtrl)

	daysOfTheMonthParser := NewDayOfTheMonthParser(mockCommonParser)
	assert.Equal(t, 1, daysOfTheMonthParser.MinAllowedValue())
}

func TestDaysOfTheMonthParser_Parse_ShouldReturnErrorWhenCommonParserReturnsAnError(t *testing.T) {
	expression := "some-expression"
	ctrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(ctrl)
	mockCommonParser.EXPECT().Parse(expression, gomock.Any()).Return([]string{}, errors.New("some error when parsing"))

	daysOfTheMonthParser := NewDayOfTheMonthParser(mockCommonParser)
	_, actualErr := daysOfTheMonthParser.Parse(expression)

	expectedErr := errors.New("some error when parsing")
	assert.Equal(t, expectedErr, actualErr)
}