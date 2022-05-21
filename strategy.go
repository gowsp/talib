package talib

// A rule for strategy building. A trading rule may be composed of a combination of other rules.
type Rule func(offset uint64) bool

// a rule which is the AND combination of this rule with the provided one
func (r Rule) And(other Rule) Rule {
	return func(offset uint64) bool {
		return r(offset) && other(offset)
	}
}

// a rule which is the OR combination of this rule with the provided one
func (r Rule) Or(other Rule) Rule {
	return func(offset uint64) bool {
		return r(offset) || other(offset)
	}
}

// A trading strategy. A strategy is a pair of complementary rules. It may recommend to enter or to exit.
type Strategy interface {
	// true to recommend to enter, false otherwise
	ShouldEnter(offset uint64) bool
	// true to recommend to exit, false otherwise
	ShouldExit(offset uint64) bool
}

type strategy struct {
	entryRule Rule
	exitRule  Rule
}

func (s *strategy) ShouldEnter(offset uint64) bool {
	return s.entryRule(offset)
}
func (s *strategy) ShouldExit(offset uint64) bool {
	return s.exitRule(offset)
}

func NewStrategy(entryRule, exitRule Rule) Strategy {
	return &strategy{entryRule, exitRule}
}
