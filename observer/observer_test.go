package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {

	random := NewRandomNumberGenerator()

	o1 := &DigitObserver{random}
	o2 := &DigitObserver{random}

	random.AddObserver(o1)
	random.AddObserver(o2)

	result := random.Execute()

	for _, r := range result {
		if len(result) != 2 && r >= 50 {
			t.Errorf("Expect result to equal random int array")
		}
	}
}
