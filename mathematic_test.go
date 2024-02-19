package main

import (
	"testing"
)

func TestDivide(t *testing.T) {
	var (
		expected int = 5
	)

	t.Run("Успешное безошибочное деление натуральных чисел", func(t *testing.T) {
		//Успешное деление натуральных чисел
		result, _ := Divide(10, 2)
		if result != expected {
			t.Errorf("result != expected!\nExpected: %d, Result:%d", expected, result)
		}
	})

}
