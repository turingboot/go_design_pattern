package factory_method

import "go_design_pattern/2-simple_factory"

// IRuleConfigParserFactory 当对象的创建逻辑比较复杂，不只是简单的 new 一下就可以，而是要组合其他类对象，
// 做各种初始化操作的时候，推荐使用工厂方法模式，
// 将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂
// IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
	CreateParser() simple_factory.IRuleConfigParser
}

// yamlRuleConfigParserFactory yamlRuleConfigParser 的工厂类
type yamlRuleConfigParserFactory struct {
}

func (y *yamlRuleConfigParserFactory) CreateParser() simple_factory.IRuleConfigParser {
	return &simple_factory.YamlRuleConfigParser{}
}

// jsonRuleConfigParserFactory jsonRuleConfigParser 的工厂类
type jsonRuleConfigParserFactory struct {
}

func (j *jsonRuleConfigParserFactory) CreateParser() simple_factory.IRuleConfigParser {
	return &simple_factory.JsonRuleConfigParser{}
}

// NewIRuleConfigParserFactory 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return &jsonRuleConfigParserFactory{}
	case "yaml":
		return &yamlRuleConfigParserFactory{}
	}
	return nil
}
