package state

import (
	"strings"
)

type state interface {
	doClock(context context, hour int)
	doUse(context context)
}

var dayInstance *dayState

type dayState struct {
}

func GetDayInstance() *dayState {
	if dayInstance == nil {
		dayInstance = &dayState{}
	}
	return dayInstance
}

func (self *dayState) doClock(context context, hour int) {
	if hour < 9 || 17 <= hour {
		context.changeState(GetNightInstance())
	}
}

func (self *dayState) doUse(context context) {
	context.recordLog("day")
}

var nightInstance *nightState

type nightState struct {
}

func GetNightInstance() *nightState {
	if nightInstance == nil {
		nightInstance = &nightState{}
	}
	return nightInstance
}

func (self *nightState) doClock(context context, hour int) {
	if 9 <= hour && hour < 17 {
		context.changeState(GetDayInstance())
	}
}

func (self *nightState) doUse(context context) {
	context.recordLog("night")
}

type context interface {
	SetClock(hour int)
	changeState(state state)
	recordLog(log string)
}

type SafeFrame struct {
	State state
	logs  []string
}

func (self *SafeFrame) SetClock(hour int) {
	self.State.doClock(self, hour)
}

func (self *SafeFrame) changeState(state state) {
	self.State = state
}

func (self *SafeFrame) recordLog(log string) {
	self.logs = append(self.logs, log)
}

func (self *SafeFrame) Use() {
	self.State.doUse(self)
}

func (self *SafeFrame) GetLog() string {
	return strings.Join(self.logs, " ")
}
