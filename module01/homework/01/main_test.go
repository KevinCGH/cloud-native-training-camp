package _01

import (
	"strings"
	"testing"
)

func TestChangeArrayString(t *testing.T) {
	arr := []string{"I", "am", "stupid", "and", "weak"}
	result := changeArrayString(arr)
	t.Logf("result: %v", result)
	if strings.Join(arr, " ") != strings.Join(result, " ") {
		t.Errorf("Expect: %v\nActual: %v", arr, result)
	}
}
