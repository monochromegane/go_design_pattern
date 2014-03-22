package visitor

import (
	"strconv"
)

type visitor interface {
	visitFile(file *file) string
	visitDir(directory *directory) string
}

type element interface {
	Accept(visitor visitor) string
}

type entry interface {
	element
	getName() string
	getSize() int
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

func (self *file) Add(entry entry) {}

func (self *file) Accept(visitor visitor) string {
	return visitor.visitFile(self)
}

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

func (self *directory) Accept(visitor visitor) string {
	return visitor.visitDir(self)
}

type listVisitor struct {
	currentDir string
}

func (self *listVisitor) visitFile(file *file) string {
	return self.currentDir + "/" + file.print(file)
}

func (self *listVisitor) visitDir(directory *directory) string {
	saveDir := self.currentDir
	list := self.currentDir + "/" + directory.print(directory)
	self.currentDir += "/" + directory.getName()
	for _, dir := range directory.dir {
		list += dir.Accept(self)
	}
	self.currentDir = saveDir
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
