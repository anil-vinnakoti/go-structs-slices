package helpers

import (
	"testing"
)

func TestAddThreeNumbers(t *testing.T){
	result := addThreeNumbers(1,2,3)
	if result != 6{
		t.Errorf("expected: %d, recieved: %d", 8, result)
	}
}