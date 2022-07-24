package parser

import (
	"cron-expression-parser/mocks"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHourParser_Parse_ShouldCallTheCommonParser_Once_AndReturnCombinedValues(t *testing.T) {
	type scenario struct {
		name          string
		hour          string
		expectedValue []string
	}
	scenarios := []scenario{
		{
			name:          "single daysOfTheWeek",
			hour:          "5",
			expectedValue: []string{"5"},
		},
		{
			name:          "range of hours",
			hour:          "5-7",
			expectedValue: []string{"5", "6", "7"},
		},
		{
			name:          "range with step",
			hour:          "5-8/1",
			expectedValue: []string{"5", "6", "7", "8"},
		},
	}

	for _, scenario := range scenarios {
		ctrl := gomock.NewController(t)
		mockCommonParser := mocks.NewMockCommonParser(ctrl)
		mockCommonParser.EXPECT().Parse(scenario.hour, gomock.Any()).Return(scenario.expectedValue,nil)

		hourParser := NewHourParser(mockCommonParser)
		actualValue, err := hourParser.Parse(scenario.hour)

		assert.Nil(t, err)
		assert.Equal(t, scenario.expectedValue, actualValue)
	}
}

func TestHourParser_Parse_ShouldCallTheCommonParserThrice_AsThereAreThreeListValues(t *testing.T) {
	hour := "3-4,6,7"
	ctrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(ctrl)
	mockCommonParser.EXPECT().Parse("3-4", gomock.Any()).Return([]string{"3", "4"},nil)
	mockCommonParser.EXPECT().Parse("6", gomock.Any()).Return([]string{"6"},nil)
	mockCommonParser.EXPECT().Parse("7", gomock.Any()).Return([]string{"7"},nil)

	hourParser := NewHourParser(mockCommonParser)
	actualValue, err := hourParser.Parse(hour)

	expectedValue := []string{"3", "4", "6", "7"}
	assert.Equal(t, expectedValue, actualValue)
	assert.Nil(t, err)
}

func TestHourParser_MaxAllowedValue_ShouldReturn_59(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(mockCtrl)

	hourParser := NewHourParser(mockCommonParser)
	assert.Equal(t, 23, hourParser.MaxAllowedValue())
}

func TestHourParser_MinAllowedValue_ShouldReturn_0(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(mockCtrl)

	hourParser := NewHourParser(mockCommonParser)
	assert.Equal(t, 0, hourParser.MinAllowedValue())
}

func TestHourParser_Parse_ShouldReturnErrorWhenCommonParserReturnsAnError(t *testing.T) {
	expression := "some-expression"
	ctrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(ctrl)
	mockCommonParser.EXPECT().Parse(expression, gomock.Any()).Return([]string{}, errors.New("some error when parsing"))

	hourParser := NewHourParser(mockCommonParser)
	_, actualErr := hourParser.Parse(expression)

	expectedErr := errors.New("some error when parsing")
	assert.Equal(t, expectedErr, actualErr)
}