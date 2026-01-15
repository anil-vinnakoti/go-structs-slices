package helpers_test

import (
	"testing"

	"github.com/anil_vinnakoti/testing/helpers"
)

func TestAddTwoNumbers(t *testing.T) {
	result := helpers.AddTwoNumbers(2, 6)
	if result != 8 {
		t.Errorf("expected: %v, recieved: %v", 8, result)
	}
}

func TestSubtractTwoNumbers(t *testing.T) {
	result := helpers.SubtractTwoNumbers(8, 5)
	if result != 3 {
		t.Errorf("expected: %v, recieved: %v", 3, result)
	}
}

func TestMultiplyTwoNumbers(t *testing.T) {
	testCases := []struct {
		testName string
		num1     int
		num2     int
		expected int
	}{
		{testName: "multiplication of 2 and 5 is 10", num1: 2, num2: 5, expected: 10},
		{testName: "multiplication of 5 and 5 is 25", num1: 5, num2: 5, expected: 25},
		{testName: "multiplication of 2 and 8 is 16", num1: 2, num2: 8, expected: 16},
		{testName: "multiplication of 8 and 5 is 40", num1: 8, num2: 5, expected: 40},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			res := helpers.MultiplyTwoNumbers(testCase.num1, testCase.num2)
			if res != testCase.expected {
				t.Errorf("expected: %d, recieved: %d", testCase.expected, res)
			}
		})
	}
}
