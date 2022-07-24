package parser_interface

type Parser interface {
	IsValid(expression string) bool
	Parse(expression string) ([]string, error)
	MinAllowedValue() int
	MaxAllowedValue() int
}
