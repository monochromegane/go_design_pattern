package memento

type memento struct {
	money int
}

func (self *memento) getMoney() int {
	return self.money
}

type Game struct {
	Money int
}

func (self *Game) CreateMemento() *memento {
	return &memento{
		self.Money,
	}
}

func (self *Game) RestoreMemento(memento *memento) {
	self.Money = memento.getMoney()
}
