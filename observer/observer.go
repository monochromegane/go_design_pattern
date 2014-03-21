package observer

import (
	"math/rand"
)

type numberGenerator struct {
	observers []observer
}

func (self *numberGenerator) addObserver(observer observer) {
	self.observers = append(self.observers, observer)
}

func (self *numberGenerator) notifyObservers() []int {
	var result []int
	for _, observer := range self.observers {
		result = append(result, observer.update())
	}
	return result
}

type randomNumberGenerator struct {
	*numberGenerator
}

type number interface {
	getNumber() int
}

func (self *randomNumberGenerator) getNumber() int {
	return rand.Intn(50)
}

func (self *randomNumberGenerator) execute() []int {
	return self.notifyObservers()
}

type observer interface {
	update() int
}

type digitObserver struct {
	generator number
}

func (self *digitObserver) update() int {
	return self.generator.getNumber()
}
