package talib

import (
	"fmt"
	"testing"
)

func TestRule(t *testing.T) {
	var rule_1 Rule = func(offset uint64) bool {
		return true
	}
	rule_2 := func(offset uint64) bool {
		return false
	}
	fmt.Println(rule_1.Or(rule_2)(0))
	fmt.Println(rule_1.And(rule_2)(0))
}
