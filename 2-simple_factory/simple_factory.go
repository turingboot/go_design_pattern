package simple_factory

import "fmt"

// IRuleConfigParser 由于 Go 本身是没有构造函数的，
// 一般而言我们采用 NewName 的方式创建对象/接口，
// 当它返回的是接口的时候，其实就是简单工厂模式
type IRuleConfigParser interface {
	Parser(data []byte)
}

type JsonRuleConfigParser struct {
}

func (j *JsonRuleConfigParser) Parser(data []byte) {
	fmt.Println("-----jsonRuleConfigParser-----")
	fmt.Println(string(data))
}

type YamlRuleConfigParser struct {
}

func (j *YamlRuleConfigParser) Parser(data []byte) {
	fmt.Println("-----yamlRuleConfigParser-----")
	fmt.Println(string(data))
}

func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return &JsonRuleConfigParser{}
	case "yaml":
		return &YamlRuleConfigParser{}
	}
	return nil
}
