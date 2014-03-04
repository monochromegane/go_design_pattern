package adapter

import (
	"testing"
)

func TestAdapterByEmbedded(t *testing.T) {
	var decorator Decorator

	decorator = NewDecorateBanner("A")

	if str := decorator.decorate(); str != "*A*" {
		t.Errorf("Expect decorated str to %s, but %s", "*A*", str)
	}

}
