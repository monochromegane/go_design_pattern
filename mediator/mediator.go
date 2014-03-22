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
	Enabled  bool
	mediator mediator
}

func (self *button) setMediator(mediator mediator) {
	self.mediator = mediator
}

func (self *button) setColleagueEnabled(enabled bool) {
	self.Enabled = enabled
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

func (self *radioButton) Check(checked bool) {
	self.checked = checked
	self.mediator.colleagueChanged()
}

type loginForm struct {
	RadioButton radioButton
	Button      button
}

func NewLoginForm() *loginForm {
	loginForm := &loginForm{}
	loginForm.createColleagues()
	return loginForm
}

func (self *loginForm) createColleagues() {
	self.RadioButton = radioButton{}
	self.Button = button{}
	self.RadioButton.setMediator(self)
	self.Button.setMediator(self)
}

func (self *loginForm) colleagueChanged() {
	if !self.RadioButton.checked {
		self.Button.setColleagueEnabled(false)
	} else {
		self.Button.setColleagueEnabled(true)
	}
}
