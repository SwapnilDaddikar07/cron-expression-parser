package parser

import (
	"cron-expression-parser/mocks"
	"cron-expression-parser/model"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCronExpressionParser_BuildRawCronExpressionComponents_ShouldParseInputInto_RawCronExpressionComponents(t *testing.T) {
	type scenario struct {
		name          string
		input         string
		expectedValue model.RawInput
	}

	scenarios := []scenario{
		{
			name:  "a random valid cron expression with 5 spaces, combination of range and step in days of the month component ",
			input: "1 2 3-4/2 4 4",
			expectedValue: model.RawInput{
				Minute:        "1",
				Hour:          "2",
				DayOfTheMonth: "3-4/2",
				Month:         "4",
				DayOfTheWeek:  "4",
			},
		},
		{
			name:  "a random valid cron expression with 5 spaces, combination of range and step in days of the month component and minute component ",
			input: "1-2 2 3-4/2 4 4",
			expectedValue: model.RawInput{
				Minute:        "1-2",
				Hour:          "2",
				DayOfTheMonth: "3-4/2",
				Month:         "4",
				DayOfTheWeek:  "4",
			},
		},
		{
			name:  "a random valid cron expression with 5 spaces, combination of range and step in days of the month component and minute component and any value identifier",
			input: "1-2 */2 3-4/2 2 4",
			expectedValue: model.RawInput{
				Minute:        "1-2",
				Hour:          "*/2",
				DayOfTheMonth: "3-4/2",
				Month:         "2",
				DayOfTheWeek:  "4",
			},
		},
	}

	for _, scenario := range scenarios {
		ctrl := gomock.NewController(t)
		minuteParser := mocks.NewMockParser(ctrl)
		hourParser := mocks.NewMockParser(ctrl)
		daysOfMonthParser := mocks.NewMockParser(ctrl)
		monthParser := mocks.NewMockParser(ctrl)
		daysOfTheWeekParser := mocks.NewMockParser(ctrl)

		cronExpressionParser := NewCronExpressionParser(minuteParser, hourParser, daysOfMonthParser, monthParser, daysOfTheWeekParser)
		actualValue, err := cronExpressionParser.BuildRawCronExpressionComponents(scenario.input)

		assert.Nil(t, err)
		assert.Equal(t, scenario.expectedValue, actualValue)
	}
}

func TestCronExpressionParser_Parse_CallRelevantParsersAnd_ReturnParsedCronExpression(t *testing.T) {
	expression := "1-2 3 4-10/2 2 3-5"
	expectedValue := model.ParsedCronExpression{
		Minutes:        []string{"1", "2"},
		Hours:          []string{"3"},
		DaysOfTheMonth: []string{"4", "6", "8", "10"},
		Months:         []string{"2"},
		DaysOfTheWeek:  []string{"3", "4", "5"},
	}

	ctrl := gomock.NewController(t)
	minuteParser := mocks.NewMockParser(ctrl)
	hourParser := mocks.NewMockParser(ctrl)
	daysOfTheMonthParser := mocks.NewMockParser(ctrl)
	monthParser := mocks.NewMockParser(ctrl)
	daysOfTheWeekParser := mocks.NewMockParser(ctrl)

	minuteParser.EXPECT().Parse("1-2").Return([]string{"1", "2"}, nil)
	hourParser.EXPECT().Parse("3").Return([]string{"3"}, nil)
	daysOfTheMonthParser.EXPECT().Parse("4-10/2").Return([]string{"4", "6", "8", "10"}, nil)
	monthParser.EXPECT().Parse("2").Return([]string{"2"}, nil)
	daysOfTheWeekParser.EXPECT().Parse("3-5").Return([]string{"3", "4", "5"}, nil)

	cronExpressionParser := NewCronExpressionParser(minuteParser, hourParser, daysOfTheMonthParser, monthParser, daysOfTheWeekParser)
	actualValue, err := cronExpressionParser.Parse(expression)

	assert.Nil(t, err)
	assert.Equal(t, expectedValue, actualValue)
}

func TestCronExpressionParser_Parse_ShouldReturnError_When_ExpressionDoesNotHaveFiveSubSections(t *testing.T) {
	type scenario struct {
		name       string
		expression string
	}
	scenarios := []scenario{
		{
			name:       "invalid input",
			expression: "   ",
		},
		{
			name:       "invalid input",
			expression: " 1 2 3-2",
		},
		{
			name:       "invalid input",
			expression: "1 2 3-2 7,2,2",
		},
	}
	expectedValue := errors.New("invalid cron expression")

	for _, scenario := range scenarios {
		ctrl := gomock.NewController(t)
		minuteParser := mocks.NewMockParser(ctrl)
		hourParser := mocks.NewMockParser(ctrl)
		dayOfTheMonthParser := mocks.NewMockParser(ctrl)
		monthParser := mocks.NewMockParser(ctrl)
		dayOfTheWeekParser := mocks.NewMockParser(ctrl)

		cronExpressionParser := NewCronExpressionParser(minuteParser, hourParser, dayOfTheMonthParser, monthParser, dayOfTheWeekParser)
		_, err := cronExpressionParser.Parse(scenario.expression)

		assert.Equal(t, expectedValue, err)
	}
}
