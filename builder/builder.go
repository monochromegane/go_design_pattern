package builder

type builder interface {
	makeTitle(title string) string
	makeString(str string) string
	makeItems(items []string) string
	close() string
}

type director struct {
	builder builder
}

func (self *director) construct() string {
	result := self.builder.makeTitle("Title")
	result += self.builder.makeString("String")
	result += self.builder.makeItems([]string{
		"Item1",
		"Item2",
	})
	result += self.builder.close()
	return result
}

type textBuilder struct {
}

func (self *textBuilder) makeTitle(title string) string {
	return "# " + title + "\n"
}

func (self *textBuilder) makeString(str string) string {
	return "## " + str + "\n"
}

func (self *textBuilder) makeItems(items []string) string {
	var result string
	for _, item := range items {
		result += "- " + item + "\n"
	}
	return result
}

func (self *textBuilder) close() string {
	return "\n"
}
