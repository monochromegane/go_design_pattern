package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	proxy := printerProxy{name: "A"}
	name := proxy.getPrinterName()

	if name != "A" {
		t.Errorf("Expect name to equal %s, but %s.\n", "A", name)
	}

	proxy.setPrinterName("B")
	name = proxy.getPrinterName()
	if name != "B" {
		t.Errorf("Expect name to equal %s, but %s.\n", "B", name)
	}

	result := proxy.print("C")
	if result != "B:C" {
		t.Errorf("Expect result to equal %s, but %s.\n", "B:C", result)
	}

}
