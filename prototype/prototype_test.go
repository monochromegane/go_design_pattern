package prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	product := Product{"A"}
	product.SetUp()

	manager := Manager{}
	manager.Register(&product)

	cloned := manager.Create("A")

	if cloned.GetName() != product.GetName() {
		t.Errorf("Expect instance to equal, but not equal.")
	}

}
