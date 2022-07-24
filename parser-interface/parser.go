package parser_interface

type Parser interface {
	Parse(expression string) ([]string, error)
	MinAllowedValue() int
	MaxAllowedValue() int
}
