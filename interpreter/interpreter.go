package interpreter

import (
	"strings"
)

type node interface {
	Parse(context *context)
	ToString() string
}

type ProgramNode struct {
	commandListNode node
}

func (self *ProgramNode) Parse(context *context) {
	context.skipToken("program")
	self.commandListNode = &commandListNode{}
	self.commandListNode.Parse(context)
}

func (self *ProgramNode) ToString() string {
	return "program: " + self.commandListNode.ToString()
}

type commandListNode struct {
	list []node
}

func (self *commandListNode) Parse(context *context) {
	for {
		if context.currentToken == "end" {
			break
		} else {
			commandNode := &commandNode{}
			commandNode.Parse(context)
			self.list = append(self.list, commandNode)
		}
	}
}

func (self *commandListNode) ToString() string {
	var str string
	for _, l := range self.list {
		str += l.ToString()
	}
	return str
}

type commandNode struct {
	node node
}

func (self *commandNode) Parse(context *context) {
	self.node = &primitiveCommandNode{}
	self.node.Parse(context)
}

func (self *commandNode) ToString() string {
	return self.node.ToString()
}

type primitiveCommandNode struct {
	name string
}

func (self *primitiveCommandNode) Parse(context *context) {
	self.name = context.currentToken
	context.skipToken(self.name)
}

func (self *primitiveCommandNode) ToString() string {
	return self.name + " "
}

type context struct {
	text         string
	tokens       []string
	currentToken string
}

func NewContext(text string) *context {
	context := &context{
		text:   text,
		tokens: strings.Fields(text),
	}
	context.nextToken()
	return context
}

func (self *context) nextToken() string {
	if len(self.tokens) == 0 {
		self.currentToken = ""
	} else {
		self.currentToken = self.tokens[0]
		self.tokens = self.tokens[1:]
	}
	return self.currentToken
}

func (self *context) skipToken(token string) {
	self.nextToken()
}
