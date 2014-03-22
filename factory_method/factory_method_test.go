package factory_method

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {

	assert := []string{"A", "B", "C"}

	factory := &IDCardFactory{}
	products := []User{
		factory.Create(factory, "A"),
		factory.Create(factory, "B"),
		factory.Create(factory, "C"),
	}

	for i, product := range products {
		if owner := product.Use(); owner != assert[i] {
			t.Errorf("Expect owner to %s, but %s.\n", assert[i], owner)
		}
	}

}
