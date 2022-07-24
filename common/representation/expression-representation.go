package representation

import "strconv"

type ExpressionRepresentation struct {
	start                 int
	end                   int
	interval              int
	isStepValueApplicable bool
}

func NewCronExpressionRepresentation() *ExpressionRepresentation {
	return &ExpressionRepresentation{}
}

func (c *ExpressionRepresentation) SetStart(start int) {
	c.start = start
}

func (c *ExpressionRepresentation) SetEnd(end int) {
	c.end = end
}

func (c *ExpressionRepresentation) SetInterval(interval int) {
	c.interval = interval
}

func (c *ExpressionRepresentation) SetIsStepApplicable(isStepApplicable bool) {
	c.isStepValueApplicable = isStepApplicable
}

func (c *ExpressionRepresentation) StepValueApplicable() bool {
	return c.isStepValueApplicable
}

func (c *ExpressionRepresentation) Execute() []string {
	result := make([]string, 0)
	for i := c.start; i <= c.end; i = i + c.interval {
		result = append(result, strconv.Itoa(i))
	}
	return result
}
