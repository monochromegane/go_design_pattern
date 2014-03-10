package abstract_factory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {

	factory := mdFactory{}

	usYahoo := factory.createLink("Yahoo!", "http://www.yahoo.com")
	jaYahoo := factory.createLink("Yahoo!Japan", "http://www.yahoo.co.jp")

	tray := factory.createTray("Yahoo!")
	tray.addToTray(usYahoo)
	tray.addToTray(jaYahoo)

	page := factory.createPage("Title", "Author")
	page.addToContent(tray)

	output := page.output()

	expect := "title: Title\nauthor: Author\n- Yahoo!\n[Yahoo!](http://www.yahoo.com)\n[Yahoo!Japan](http://www.yahoo.co.jp)\n\n"

	if output != expect {
		t.Errorf("Expect output to %s, but %s\n", expect, output)
	}

}
