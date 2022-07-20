package simple

import (
	"testing"
)

func TestSum(t *testing.T) {
	expected := 3
	num1, num2 := 1, 2

	if Plus(num1, num2) != expected {
		t.Errorf("Plus incorrect")
	}
}

func TestMinus(t *testing.T) {
	expected := 1
	num1, num2 := 2, 1

	if Minus(num1, num2) != expected {
		t.Errorf("Minus incorrect")
	}
}
