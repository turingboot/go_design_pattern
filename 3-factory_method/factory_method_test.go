package factory_method

import (
	"testing"
)

func TestNewIRuleConfigParserFactory(t *testing.T) {
	j := NewIRuleConfigParserFactory("json")
	j.CreateParser().Parser([]byte("hello json"))

	y := NewIRuleConfigParserFactory("yaml")
	y.CreateParser().Parser([]byte("hello yaml"))
}
