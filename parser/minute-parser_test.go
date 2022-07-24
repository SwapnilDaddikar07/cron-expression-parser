package parser

import (
	"cron-expression-parser/mocks"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinuteParser_Parse_ShouldCallTheCommonParser_Once_WithGivenValues(t *testing.T) {
	type scenario struct {
		name          string
		minute        string
		expectedValue []string
	}
	scenarios := []scenario{
		{
			name:          "single minute",
			minute:        "1",
			expectedValue: []string{"1"},
		},
		{
			name:          "range of minutes",
			minute:        "2-5",
			expectedValue: []string{"2", "3", "4", "5"},
		},
		{
			name:          "range with step",
			minute:        "3-7/2",
			expectedValue: []string{"3", "5", "7"},
		},
	}
	for _, scenario := range scenarios {
		mockCtrl := gomock.NewController(t)
		mockCommonParser := mocks.NewMockCommonParser(mockCtrl)
		mockCommonParser.EXPECT().Parse(scenario.minute, gomock.Any()).Return(scenario.expectedValue, nil)

		minuteParser := NewMinuteParser(mockCommonParser)
		actualValue, err := minuteParser.Parse(scenario.minute)

		assert.Equal(t, scenario.expectedValue, actualValue)
		assert.Nil(t, err)
	}
}

func TestMinuteParser_Parse_ShouldCallTheCommonParser_Thrice_AsThereAreThreeListSeparatedValues_AndReturnCombinedResult(t *testing.T) {
	minute := "4,5,6"

	mockCtrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(mockCtrl)
	mockCommonParser.EXPECT().Parse("4", gomock.Any()).Return([]string{"4"}, nil).Times(1)
	mockCommonParser.EXPECT().Parse("5", gomock.Any()).Return([]string{"5"}, nil).Times(1)
	mockCommonParser.EXPECT().Parse("6", gomock.Any()).Return([]string{"6"}, nil).Times(1)

	minuteParser := NewMinuteParser(mockCommonParser)
	actualValue, err := minuteParser.Parse(minute)

	expectedValue := []string{"4", "5", "6"}
	assert.Equal(t, expectedValue, actualValue)
	assert.Nil(t, err)
}

func TestMinuteParser_Parse_ShouldCallTheParserTwice_AsThereAreTwoListSeparatedValues_AndReturnCombinedResult(t *testing.T) {
	minute := "2,3-7/2"
	mockCtrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(mockCtrl)
	mockCommonParser.EXPECT().Parse("2", gomock.Any()).Return([]string{"2"}, nil)
	mockCommonParser.EXPECT().Parse("3-7/2", gomock.Any()).Return([]string{"3", "5", "7"}, nil)

	minuteParser := NewMinuteParser(mockCommonParser)
	actualValue, err := minuteParser.Parse(minute)

	expectedValue := []string{"2", "3", "5", "7"}
	assert.Equal(t, expectedValue, actualValue)
	assert.Nil(t, err)
}

func TestMinuteParser_MaxAllowedValue_ShouldReturn_59(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(ctrl)

	minuteParser := NewMinuteParser(mockCommonParser)
	assert.Equal(t, 59, minuteParser.MaxAllowedValue())
}

func TestMinuteParser_MinAllowedValue_ShouldReturn_0(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(ctrl)

	minuteParser := NewMinuteParser(mockCommonParser)
	assert.Equal(t, 0, minuteParser.MinAllowedValue())
}

func TestMinuteParser_Parse_ShouldReturnErrorWhenCommonParserReturnsAnError(t *testing.T) {
	expression := "some-expression"
	ctrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(ctrl)
	mockCommonParser.EXPECT().Parse(expression, gomock.Any()).Return([]string{}, errors.New("some error when parsing"))

	minuteParser := NewMinuteParser(mockCommonParser)
	_, actualErr := minuteParser.Parse(expression)

	expectedErr := errors.New("some error when parsing")
	assert.Equal(t, expectedErr, actualErr)
}
