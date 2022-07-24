package main

import (
	"cron-expression-parser/common/parser"
	parser2 "cron-expression-parser/parser"
	console_writer "cron-expression-parser/writer/console-writer"
	"fmt"
	"os"
)

func main() {
	commonParser := parser.NewCommonParser()
	minuteParser := parser2.NewMinuteParser(commonParser)
	hourParser := parser2.NewHourParser(commonParser)
	daysOfTheMonthParser := parser2.NewDayOfTheMonthParser(commonParser)
	monthParser := parser2.NewMonthParser(commonParser)
	daysOfTheWeekParser := parser2.NewDaysOfTheWeekParser(commonParser)

	cronExpressionParser := parser2.NewCronExpressionParser(minuteParser, hourParser, daysOfTheMonthParser, monthParser, daysOfTheWeekParser)

	consoleWriter := console_writer.NewConsoleWriter()

	fmt.Println("Enter a cron expression")

	for _, input := range os.Args[1:] {
		parsedCronJobExpression, err := cronExpressionParser.Parse(input)
		if err != nil {
			fmt.Printf("error occured when parsing expression %v \n", err)
			fmt.Println("Enter a cron expression")
			continue
		}
		err = consoleWriter.Write(parsedCronJobExpression)
		if err != nil {
			fmt.Printf("error occur when writing to console %v \n", err)
		}
		fmt.Println("Enter a cron expression")
	}
}
