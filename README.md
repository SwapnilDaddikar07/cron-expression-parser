# cron-expression-parser

A cron expression is commonly used with a cron job.
A cron job runs if the current date/time satisfies the expression.

Other resources on cron expressions
Cron expression generator - https://crontab.guru/
Cron expression - https://en.wikipedia.org/wiki/Cron

This program takes in a cron expression and outputs the parsed values in a tabular format.

Input , output and tabular format can be found here -
https://drive.google.com/file/d/1QmPHIZmHId7Shy-kpA0MmjJV1aDOY15s/view

As per the problem statement , one of the constraints is as follows - 
**The output should be formatted as a table with the field name taking the first 14 columns and the times as a space-separated list following it.**

Additional libraries used
**gomock**  - https://github.com/golang/mock (A library for creating mocks. Helps in unit testing)
**testify** - https://github.com/stretchr/testify (A library to gracefully handle assertions in unit tests.)

**Codebase**

**How to run the code ?**

The project contains a Makefile.

**make install_deps**
This downloads the mentioned additional libraries.

**make console_reader**
Starts a scanner which takes input from the console and prints out the output in the required format.

**make run_preloaded_values**
To run and check output for some preloaded valid cron expressions. Check the makefile to see the expressions that will run.


**Codebase context**

Cron Expression Parser is the entry point to the parser.
It exposes a **Parse** method which takes in an expression and returned a parsed value.
The parsed value is represented by a model named ParsedCronExpression.

There are separate parsers for each component of the cron expression i.e minute/hour/month etc.
These parses hold specific information pertaining to their corresponding component.
Every component can be parsed using the same algorithm and to support the reuse of this algorithm , the codebase contains a common parser named **CommonParser**

To print the output to console , there is a console-writer struct which internally knows the format to be used for printing.

Code flow sequence
CronExpressionParser --> IndividualParsers(minute/hour/day etc) --> CommonParser --> ConsoleWriter

**Algorithm**

1. First preference given to *.
   If the component contains just a * (any identifier) , the code fetches the min and max value from respective parser and collect all values.

2. An expression which contains a "/" i.e (step identifier) will always be accompanied by another identifier.
   Example [2-3/2] here **/2** is the step identifier accompanied by a range identifier.
   [*/2] here **/2** is the step identifier accompanied by a * (any identifier)

3. So if a step identifier is present , the code parses the expression and remember the step value.

4. If the expression contains a range identifier , the code simply splits the expression at the range identifier **"-"**.
   There are two possibilities now. If the step identifier is present , the program utilizes the step identifier for incrementing values.
   If not , the increment factor becomes 1.

5. If the expression contains a split identifier , the program simply split the input at the separator **","**.
   There are two possibilities. If a step identifier was encountered previously , the last value of the split identifier will serve as the starting point.
   The end value of the expression will be found out by calling the respective parser. Increment from start to end by the step identifier.
   If step identifier is not present , the code will simply include comma separated values in the output.

6. Finally , if there is no split identifier or a list identifier , it means there is a singular value.
   The singular value can either be any identifier * or just a number.
   If the single value is just a number and step value was not encountered previously , the code will return the single number back.
   If a step value was encountered previously and the singular value is * , the program will find the starting point and ending range by calling respective parser.
   The step value will be used as the incrementing factor.
   If the single value is a number and step value was encountered , the number will be used as starting point. The ending range will be found by calling respective parser.


