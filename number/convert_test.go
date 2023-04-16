package number

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

func NumberConverterTest(t *testing.T) {
	rule := new(MockReplaceRuleInterface)
	rule.On("Replace", 1).Return("Replaced")
	nc := NewNumberConverter([]ReplaceRuleInterface{rule})
	if nc.Convert(1) != "Replaced" {
		t.Errorf("Convert(1) = %s; want \"Replaced\"", nc.Convert(1))
	}
}

type MockReplaceRuleInterface struct {
	mock.Mock
}

func (m *MockReplaceRuleInterface) Replace(n int) string {
	args := m.Called(n)
	return args.String(0)
}

func (t *MockReplaceRuleInterface) TestConvertWithUnmatchedFizzBuzzRulesAndConstantRule() {
	fizzBuzz := NewNumberConverter([]ReplaceRuleInterface{
			t.createMockRule(1, ""),
			t.createMockRule(1, ""),
			t.createMockRule(1, "1"),
	})
	t.assertEqual("1", fizzBuzz.Convert(1))
}

func (t *MockReplaceRuleInterface) createMockRule(expectedNumber int, replacement string) ReplaceRuleInterface {
	rule := &MockReplaceRule{}
	rule.On("Replace", expectedNumber).Return(replacement)
	return rule
}

// func (nc *NumberConverterTest) createMockRule(expectedNumber int, replacement string) ReplaceRuleInterface {
// 	rule := new(MockReplaceRule)
// 	rule.On("replace", expectedNumber).Return(replacement).Once()
// 	return rule
// }

// func (nc *NumberConverterTest) TestConvertWithFizzBuzzRules() {
// 	fizzRule := nc.createMockRule(1, "Fizz")
// 	buzzRule := nc.createMockRule(1, "Buzz")
// 	fBRule := nc.createMockRule(1, "FizzBuzz")
// 	fBConverter := NewNumberConverter([]ReplaceRuleInterface{fizzRule, buzzRule, fBRule})

// 	nc.Assert().Equal("FizzBuzz", fBConverter.Convert(1))
// }