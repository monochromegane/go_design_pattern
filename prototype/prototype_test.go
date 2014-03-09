package prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	product := product{"A"}
	product.setUp()

	manager := manager{}
	manager.register(&product)

	cloned := manager.create("A")

	if cloned.getName() != product.getName() {
		t.Errorf("Expect instance to equal, but not equal.")
	}

}
