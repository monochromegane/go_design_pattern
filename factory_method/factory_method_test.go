package factory_method

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {

	assert := []string{"A", "B", "C"}

	factory := &IDCardFactory{}
	products := []user{
		factory.create(factory, "A"),
		factory.create(factory, "B"),
		factory.create(factory, "C"),
	}

	for i, product := range products {
		if owner := product.use(); owner != assert[i] {
			t.Errorf("Expect owner to %s, but %s.\n", assert[i], owner)
		}
	}

}
