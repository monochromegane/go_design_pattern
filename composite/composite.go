package composite

import (
	"strconv"
)

type entry interface {
	getName() string
	getSize() int
	PrintList(prefix string) string
	Add(entry entry)
}

type defaultEntry struct {
	entry
	name string
}

func (self *defaultEntry) getName() string {
	return self.name
}

func (self *defaultEntry) print(entry entry) string {
	return entry.getName() + " (" + strconv.Itoa(entry.getSize()) + ")\n"
}

type file struct {
	*defaultEntry
	size int
}

func (self *file) getSize() int {
	return self.size
}

func (self *file) PrintList(prefix string) string {
	return prefix + "/" + self.print(self)
}

func (self *file) Add(entry entry) {}

type directory struct {
	*defaultEntry
	dir []entry
}

func (self *directory) getSize() int {
	size := 0
	for _, dir := range self.dir {
		size += dir.getSize()
	}
	return size
}

func (self *directory) Add(entry entry) {
	self.dir = append(self.dir, entry)
}

func (self *directory) PrintList(prefix string) string {
	list := prefix + "/" + self.print(self)
	for _, dir := range self.dir {
		list += dir.PrintList(prefix + "/" + self.getName())
	}
	return list
}

// 利用側に埋込構造体を意識させないためのインスタンス生成関数。
func NewFile(name string, size int) *file {
	return &file{
		defaultEntry: &defaultEntry{name: name}, // 埋込時のキーには構造体と同名のものを使うことができる
		size:         size,
	}
}
func NewDirectory(name string) *directory {
	return &directory{defaultEntry: &defaultEntry{name: name}}
}
