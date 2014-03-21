package memento

type memento struct {
	money int
}

func (self *memento) getMoney() int {
	return self.money
}

type game struct {
	money int
}

func (self *game) createMemento() *memento {
	return &memento{
		self.money,
	}
}

func (self *game) restoreMemento(memento *memento) {
	self.money = memento.getMoney()
}
