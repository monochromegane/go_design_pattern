package bridge

type DisplayImpl interface {
	rawOpen() string
	rawPrint() string
	rawClose() string
}

type StringDisplayImpl struct {
	str string
}

func (self *StringDisplayImpl) rawOpen() string {
	return self.printLine()
}

func (self *StringDisplayImpl) rawPrint() string {
	return "|" + self.str + "|\n"
}

func (self *StringDisplayImpl) rawClose() string {
	return self.printLine()
}

func (self *StringDisplayImpl) printLine() string {
	str := "+"
	for _, _ = range self.str {
		str += string("-")
	}
	str += "+\n"
	return str
}

type Display struct {
	impl DisplayImpl
}

func (self *Display) open() string {
	return self.impl.rawOpen()
}

func (self *Display) print() string {
	return self.impl.rawPrint()
}

func (self *Display) close() string {
	return self.impl.rawClose()
}

func (self *Display) display() string {
	return self.open() +
		self.print() +
		self.close()
}

type CountDisplay struct {
	*Display
}

func (self *CountDisplay) multiDisplay(num int) string {
	str := self.open()
	for i := 0; i < num; i++ {
		str += self.print()
	}
	str += self.close()
	return str
}
