package proxy

type printable interface {
	setPrinterName(name string)
	getPrinterName() string
	print(str string) string
}

type printer struct {
	name string
}

func (self *printer) setPrinterName(name string) {
	self.name = name
}

func (self *printer) getPrinterName() string {
	return self.name
}

func (self *printer) print(str string) string {
	return self.name + ":" + str
}

type printerProxy struct {
	name string
	real *printer
}

func (self *printerProxy) setPrinterName(name string) {
	if self.real != nil {
		self.real.setPrinterName(name)
	}
	self.name = name
}

func (self *printerProxy) getPrinterName() string {
	return self.name
}

func (self *printerProxy) print(str string) string {
	self.realize()
	return self.real.print(str)
}

func (self *printerProxy) realize() {
	if self.real == nil {
		self.real = &printer{self.name}
	}
}
