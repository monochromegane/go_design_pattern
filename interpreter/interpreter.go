package interpreter

import (
	"strings"
)

type node interface {
	parse(context *context)
	toString() string
}

type programNode struct {
	commandListNode node
}

func (self *programNode) parse(context *context) {
	context.skipToken("program")
	self.commandListNode = &commandListNode{}
	self.commandListNode.parse(context)
}

func (self *programNode) toString() string {
	return "program: " + self.commandListNode.toString()
}

type commandListNode struct {
	list []node
}

func (self *commandListNode) parse(context *context) {
	for {
		if context.currentToken == "end" {
			break
		} else {
			commandNode := &commandNode{}
			commandNode.parse(context)
			self.list = append(self.list, commandNode)
		}
	}
}

func (self *commandListNode) toString() string {
	var str string
	for _, l := range self.list {
		str += l.toString()
	}
	return str
}

type commandNode struct {
	node node
}

func (self *commandNode) parse(context *context) {
	self.node = &primitiveCommandNode{}
	self.node.parse(context)
}

func (self *commandNode) toString() string {
	return self.node.toString()
}

type primitiveCommandNode struct {
	name string
}

func (self *primitiveCommandNode) parse(context *context) {
	self.name = context.currentToken
	context.skipToken(self.name)
}

func (self *primitiveCommandNode) toString() string {
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
