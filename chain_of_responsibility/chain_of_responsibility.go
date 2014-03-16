package chain_of_responsibility

import (
	"strconv"
)

type trouble struct {
	number int
}

func (self *trouble) getNumber() int {
	return self.number
}

type support interface {
	resolve(trouble trouble) bool
	handle(support support, trouble trouble) string
}

type baseSupport struct {
	support
	name string
	next support
}

func (self *baseSupport) setNext(next support) {
	self.next = next
}

func (self *baseSupport) handle(support support, trouble trouble) string {
	if support.resolve(trouble) {
		return self.done(trouble)
	} else if self.next != nil {
		return self.next.handle(self.next, trouble)
	} else {
		return self.fail(trouble)
	}
}

func (self *baseSupport) done(trouble trouble) string {
	return "trouble:" + strconv.Itoa(trouble.getNumber()) + " is resolved by " + self.name
}

func (self *baseSupport) fail(trouble trouble) string {
	return "trouble:" + strconv.Itoa(trouble.getNumber()) + " cannot be resolved"
}

type noSupport struct {
	*baseSupport
}

func (self *noSupport) resolve(trouble trouble) bool {
	return false
}

func newNoSupport(name string) *noSupport {
	return &noSupport{&baseSupport{name: name}}
}

type limitSupport struct {
	*baseSupport
	limit int
}

func (self *limitSupport) resolve(trouble trouble) bool {
	if trouble.getNumber() < self.limit {
		return true
	} else {
		return false
	}
}

func newLimitSupport(name string, limit int) *limitSupport {
	return &limitSupport{&baseSupport{name: name}, limit}
}
