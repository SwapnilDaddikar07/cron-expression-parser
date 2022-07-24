package parser

import (
	"cron-expression-parser/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMonthParser_Parse_ShouldCallTheCommonParser_Once_WithGivenInput(t *testing.T) {
	type scenario struct {
		name        string
		month       string
		expectedVal []string
	}
	scenarios := []scenario{
		{
			name:        "single month",
			month:       "2",
			expectedVal: []string{"2"},
		},
		{
			name:        "range of months",
			month:       "3-4",
			expectedVal: []string{"3", "4"},
		},
		{
			name:        "range with step",
			month:       "3-7/2",
			expectedVal: []string{"3", "5", "7"},
		},
	}

	for _, scenario := range scenarios {
		mockCtrl := gomock.NewController(t)
		mockCommonParser := mocks.NewMockCommonParser(mockCtrl)
		mockCommonParser.EXPECT().Parse(scenario.month, gomock.Any()).Return(scenario.expectedVal,nil).Times(1)

		monthParser := NewMonthParser(mockCommonParser)
		actualValue, err := monthParser.Parse(scenario.month)

		assert.Equal(t, scenario.expectedVal, actualValue)
		assert.Nil(t, err)

	}

}

func TestMonthParser_Parse_ShouldCallTheCommonParser_Twice_AsThereAreTwoListValues(t *testing.T) {
	month := "3-5,6-8"

	mockCtrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(mockCtrl)
	mockCommonParser.EXPECT().Parse("3-5", gomock.Any()).Return([]string{"3", "4", "5"},nil).Times(1)
	mockCommonParser.EXPECT().Parse("6-8", gomock.Any()).Return([]string{"6", "7", "8"},nil).Times(1)

	monthParser := NewMonthParser(mockCommonParser)
	actualValue, err := monthParser.Parse(month)

	expectedValue := []string{"3", "4", "5", "6", "7", "8"}
	assert.Equal(t, expectedValue, actualValue)
	assert.Nil(t, err)
}

func TestMonthParser_Parse_ShouldCallTheCommonParser_Thrice_AsThereAreThreeListValues(t *testing.T) {
	month := "3-5,6-8,*/6"

	mockCtrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(mockCtrl)
	mockCommonParser.EXPECT().Parse("3-5", gomock.Any()).Return([]string{"3", "4", "5"},nil).Times(1)
	mockCommonParser.EXPECT().Parse("6-8", gomock.Any()).Return([]string{"6", "7", "8"},nil).Times(1)
	mockCommonParser.EXPECT().Parse("*/6", gomock.Any()).Return([]string{"6", "12"},nil).Times(1)

	monthParser := NewMonthParser(mockCommonParser)
	actualValue, err := monthParser.Parse(month)

	expectedValue := []string{"3", "4", "5", "6", "7", "8", "6", "12"}
	assert.Equal(t, expectedValue, actualValue)
	assert.Nil(t, err)
}

func TestMonthParser_MaxAllowedValue_ShouldReturn_12(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(mockCtrl)

	monthParser := NewMonthParser(mockCommonParser)

	assert.Equal(t, 12, monthParser.MaxAllowedValue())
}

func TestMonthParser_MinAllowedValue_ShouldReturn_1(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCommonParser := mocks.NewMockCommonParser(mockCtrl)

	monthParser := NewMonthParser(mockCommonParser)

	assert.Equal(t, 1, monthParser.MinAllowedValue())
}
