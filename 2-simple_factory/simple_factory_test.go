package simple_factory

import (
	"testing"
)

func TestNewIRuleConfigParser(t *testing.T) {
	j := NewIRuleConfigParser("json")
	j.Parser([]byte("hello"))

	y := NewIRuleConfigParser("yaml")
	y.Parser([]byte("hello"))

}
