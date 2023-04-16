package number

import "strings"

type ReplaceRuleInterface interface {
	Replace(n int) string
}

type NumberConverter struct {
	rules []ReplaceRuleInterface
}

func NewNumberConverter(rules []ReplaceRuleInterface) *NumberConverter {
	return &NumberConverter{
		rules,
	}
}

// func (nc *NumberConverter) Convert(n int) string {
// 	if len(nc.rules) > 0 {
// 		return nc.rules[0].replace(n)
// 	}
// 	return ""
// }

func (nc *NumberConverter) Convert(n int) string {
	var result strings.Builder
	for _, rule := range nc.rules {
			result.WriteString(rule.Replace(n))
	}
	return result.String()
}

// func (nc NumberConverter) Convert(n int) string {
// 	if n % 3 == 0 && n % 5 == 0 {
// 		return "FizzBuzz"
// 	} else if n % 3 == 0 {
// 		return "Fizz"
// 	} else if n % 5 == 0 {
// 		return "Buzz"
// 	}
// 	return strconv.Itoa(n)
// }