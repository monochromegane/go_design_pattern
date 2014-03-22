package chain_of_responsibility

import (
	"strconv"
)

type Trouble struct {
	number int
}

func (self *Trouble) getNumber() int {
	return self.number
}

type support interface {
	resolve(trouble Trouble) bool
	Handle(support support, trouble Trouble) string
}

type defaultSupport struct {
	support
	name string
	next support
}

func (self *defaultSupport) SetNext(next support) {
	self.next = next
}

func (self *defaultSupport) Handle(support support, trouble Trouble) string {
	if support.resolve(trouble) {
		return self.done(trouble)
	} else if self.next != nil {
		return self.next.Handle(self.next, trouble)
	} else {
		return self.fail(trouble)
	}
}

func (self *defaultSupport) done(trouble Trouble) string {
	return "trouble:" + strconv.Itoa(trouble.getNumber()) + " is resolved by " + self.name
}

func (self *defaultSupport) fail(trouble Trouble) string {
	return "trouble:" + strconv.Itoa(trouble.getNumber()) + " cannot be resolved"
}

type noSupport struct {
	*defaultSupport
}

func (self *noSupport) resolve(trouble Trouble) bool {
	return false
}

func NewNoSupport(name string) *noSupport {
	return &noSupport{&defaultSupport{name: name}}
}

type limitSupport struct {
	*defaultSupport
	limit int
}

func (self *limitSupport) resolve(trouble Trouble) bool {
	if trouble.getNumber() < self.limit {
		return true
	} else {
		return false
	}
}

func NewLimitSupport(name string, limit int) *limitSupport {
	return &limitSupport{&defaultSupport{name: name}, limit}
}
