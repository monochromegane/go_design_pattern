package mediator

type mediator interface {
	createColleagues()
	colleagueChanged()
}

type colleague interface {
	setMediator(mediator mediator)
	setColleagueEnabled(enabled bool)
}

type button struct {
	enabled  bool
	mediator mediator
}

func (self *button) setMediator(mediator mediator) {
	self.mediator = mediator
}

func (self *button) setColleagueEnabled(enabled bool) {
	self.enabled = enabled
}

type radioButton struct {
	enabled  bool
	checked  bool
	mediator mediator
}

func (self *radioButton) setMediator(mediator mediator) {
	self.mediator = mediator
}

func (self *radioButton) setColleagueEnabled(enabled bool) {
	self.enabled = enabled
}

func (self *radioButton) check(checked bool) {
	self.checked = checked
	self.mediator.colleagueChanged()
}

type loginForm struct {
	radioButton radioButton
	button      button
}

func newLoginForm() *loginForm {
	loginForm := &loginForm{}
	loginForm.createColleagues()
	return loginForm
}

func (self *loginForm) createColleagues() {
	self.radioButton = radioButton{}
	self.button = button{}
	self.radioButton.setMediator(self)
	self.button.setMediator(self)
}

func (self *loginForm) colleagueChanged() {
	if !self.radioButton.checked {
		self.button.setColleagueEnabled(false)
	} else {
		self.button.setColleagueEnabled(true)
	}
}
