package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {

	random := &randomNumberGenerator{&numberGenerator{}}

	o1 := &digitObserver{random}
	o2 := &digitObserver{random}

	random.addObserver(o1)
	random.addObserver(o2)

	result := random.execute()

	for _, r := range result {
		if len(result) != 2 && r >= 50 {
			t.Errorf("Expect result to equal random int array")
		}
	}
}
