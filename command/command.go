package command

import (
	"strconv"
)

type command interface {
	execute() string
}

type macroCommand struct {
	commands []command
}

func (self *macroCommand) execute() string {
	var result string
	for _, command := range self.commands {
		result += command.execute() + "\n"
	}
	return result
}

func (self *macroCommand) append(command command) {
	self.commands = append(self.commands, command)
}

func (self *macroCommand) undo() {
	if len(self.commands) != 0 {
		self.commands = self.commands[:len(self.commands)-1]
	}
}

func (self *macroCommand) clear() {
	self.commands = []command{}
}

type position struct {
	x, y int
}

type drawCommand struct {
	position *position
}

func (self *drawCommand) execute() string {
	return strconv.Itoa(self.position.x) + "." + strconv.Itoa(self.position.y)
}
