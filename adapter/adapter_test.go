package adapter

import (
	"testing"
)

func TestAdapterByEmbedded(t *testing.T) {
	var decorator Decorator

	decorator = NewEmbeddedDecorateBanner("A")

	if str := decorator.Decorate(); str != "*A*" {
		t.Errorf("Expect decorated str to %s, but %s", "*A*", str)
	}

}

func TestAdapterByDelegate(t *testing.T) {
	var decorator Decorator

	decorator = NewCompositionDecorateBanner("A")

	if str := decorator.Decorate(); str != "*A*" {
		t.Errorf("Expect decorated str to %s, but %s", "*A*", str)
	}

}
