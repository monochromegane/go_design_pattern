package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	proxy := PrinterProxy{Name: "A"}
	name := proxy.GetPrinterName()

	if name != "A" {
		t.Errorf("Expect name to equal %s, but %s.\n", "A", name)
	}

	proxy.SetPrinterName("B")
	name = proxy.GetPrinterName()
	if name != "B" {
		t.Errorf("Expect name to equal %s, but %s.\n", "B", name)
	}

	result := proxy.Print("C")
	if result != "B:C" {
		t.Errorf("Expect result to equal %s, but %s.\n", "B:C", result)
	}

}
