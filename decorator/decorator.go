package decorator

type display interface {
	getColumns() int
	getRows() int
	getRowText(row int) string
	show(display display) string
}

type baseDisplay struct {
	display
}

func (self *baseDisplay) show(display display) string {
	str := ""
	for i := 0; i < display.getRows(); i++ {
		str += display.getRowText(i)
	}
	return str
}

type stringDisplay struct {
	*baseDisplay
	str string
}

func (self *stringDisplay) getColumns() int {
	return len(self.str)
}

func (self *stringDisplay) getRows() int {
	return 1
}

func (self *stringDisplay) getRowText(row int) string {
	if row == 0 {
		return self.str
	} else {
		return ""
	}
}

type border struct {
	*baseDisplay
	display display
}

type sideBorder struct {
	*border
	borderChar string
}

func (self *sideBorder) getColumns() int {
	return len(self.borderChar)*2 + self.display.getColumns()
}

func (self *sideBorder) getRows() int {
	return self.display.getRows()
}

func (self *sideBorder) getRowText(row int) string {
	return self.borderChar + self.display.getRowText(row) + self.borderChar
}
