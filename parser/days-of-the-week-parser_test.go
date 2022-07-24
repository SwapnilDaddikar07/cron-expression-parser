package parser

import (
	"cron-expression-parser/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDaysOfTheWeekParser_Parse_ShouldCallTheCommonParser_Once_AndReturnCombinedValues(t *testing.T) {
	type scenario struct {
		name          string
		daysOfTheWeek string
		expectedValue []string
	}
	scenarios := []scenario{
		{
			name:          "single daysOfTheWeek",
			daysOfTheWeek: "5",
			expectedValue: []string{"5"},
		},
		{
			name:          "range of daysOfTheWeek",
			daysOfTheWeek: "1-2",
			expectedValue: []string{"1", "2"},
		},
		{
			name:          "range with step",
			daysOfTheWeek: "1-3/1",
			expectedValue: []string{"1", "2", "3"},
		},
	}

	for _, scenario := range scenarios {
		ctrl := gomock.NewController(t)
		mockCommonParser := mocks.NewMockCommonParser(ctrl)
		mockCommonParser.EXPECT().Parse(scenario.daysOfTheWeek, gomock.Any()).Return(scenario.expectedValue,nil)

		daysOfTheWeekParser := NewDaysOfTheWeekParser(mockCommonParser)
		actualValue, err := daysOfTheWeekParser.Parse(scenario.daysOfTheWeek)

		assert.Equal(t, scenario.expectedValue, actualValue)
		assert.Nil(t, err)
	}
}

func TestDaysOfTheWeekParser_Parse_ShouldCallTheCommonParserThrice_AsThereAreThreeListValues(t *testing.T) {
	daysOfTheWeek := "3-4,6,7"
	ctrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(ctrl)
	mockCommonParser.EXPECT().Parse("3-4", gomock.Any()).Return([]string{"3", "4"},nil)
	mockCommonParser.EXPECT().Parse("6", gomock.Any()).Return([]string{"6"},nil)
	mockCommonParser.EXPECT().Parse("7", gomock.Any()).Return([]string{"7"},nil)

	daysOfTheWeekParser := NewDaysOfTheWeekParser(mockCommonParser)
	actualValue, err := daysOfTheWeekParser.Parse(daysOfTheWeek)

	expectedValue := []string{"3", "4", "6", "7"}
	assert.Equal(t, expectedValue, actualValue)
	assert.Nil(t, err)
}

func TestDaysOfTheWeekParser_MaxAllowedValue_ShouldReturn_59(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(mockCtrl)

	daysOfTheWeekParser := NewDaysOfTheWeekParser(mockCommonParser)
	assert.Equal(t, 6, daysOfTheWeekParser.MaxAllowedValue())
}

func TestDaysOfTheWeekParser_MinAllowedValue_ShouldReturn_0(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(mockCtrl)

	daysOfTheWeekParser := NewDaysOfTheWeekParser(mockCommonParser)
	assert.Equal(t, 0, daysOfTheWeekParser.MinAllowedValue())
}
