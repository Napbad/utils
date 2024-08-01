package test

import (
	"testing"

	"github.com/Napbad/utils/check"
)

// TestCheck 测试函数
func TestCheck(t *testing.T) {
	var nilStr *string
	var emptyStr *string = new(string)
	*emptyStr = ""
	var spaceStr *string = new(string)
	*spaceStr = "   "

	tests := []struct {
		input    *string
		expected bool
	}{
		{nilStr, true},
		{emptyStr, true},
		{spaceStr, true},
		{newString("hello"), false},
	}

	for _, test := range tests {
		result, err := check.IsNilOrBlank(test.input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != test.expected {
			t.Errorf("For input '%v', expected %v but got %v", test.input, test.expected, result)
		}
	}
}

// newString 是一个辅助函数，用于创建一个指向字符串的指针
func newString(s string) *string {
	return &s
}
