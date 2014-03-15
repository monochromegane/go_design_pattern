package composite

import (
	"strconv"
)

type entry interface {
	getName() string
	getSize() int
	printList(prefix string) string
	add(entry entry)
}

type baseEntry struct {
	entry
	name string
}

func (self *baseEntry) getName() string {
	return self.name
}

func (self *baseEntry) print(entry entry) string {
	return entry.getName() + " (" + strconv.Itoa(entry.getSize()) + ")\n"
}

type file struct {
	*baseEntry
	size int
}

func (self *file) getSize() int {
	return self.size
}

func (self *file) printList(prefix string) string {
	return prefix + "/" + self.print(self)
}

func (self *file) add(entry entry) {}

type directory struct {
	*baseEntry
	dir []entry
}

func (self *directory) getSize() int {
	size := 0
	for _, dir := range self.dir {
		size += dir.getSize()
	}
	return size
}

func (self *directory) add(entry entry) {
	self.dir = append(self.dir, entry)
}

func (self *directory) printList(prefix string) string {
	list := prefix + "/" + self.print(self)
	for _, dir := range self.dir {
		list += dir.printList(prefix + "/" + self.getName())
	}
	return list
}

// 利用側に埋込構造体を意識させないためのインスタンス生成関数。
func NewFile(name string, size int) *file {
	return &file{
		baseEntry: &baseEntry{name: name}, // 埋込時のキーには構造体と同名のものを使うことができる
		size:      size,
	}
}
func NewDirectory(name string) *directory {
	return &directory{baseEntry: &baseEntry{name: name}}
}
