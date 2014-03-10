package abstract_factory

// 継承ではなくインターフェースによるポリモフィズムを用いる
type item interface {
	toString() string
}

// インターフェースの埋込も可能
type link interface {
	item
}

// インターフェースのポリモフィズムでは構造体間の共通実装やプロパティを
// 利用することができないため、インターフェースを実装した構造体を埋め込む
// ことで実現する
type tray interface {
	item
	addToTray(item item)
}

type baseTray struct {
	tray []item
}

func (self *baseTray) addToTray(item item) {
	self.tray = append(self.tray, item)
}

type page interface {
	addToContent(item item)
	output() string
}

type basePage struct {
	content []item
}

func (self *basePage) addToContent(item item) {
	self.content = append(self.content, item)
}

type factory interface {
	createLink(caption, url string) link
	createTray(caption string) tray
	createPage(title, author string) page
}

type mdLink struct {
	caption, url string
}

func (self *mdLink) toString() string {
	return "[" + self.caption + "](" + self.url + ")"
}

type mdTray struct {
	// 共通処理とプロパティを扱える構造体、かつインターフェースによる
	// ポリモフィズムも可能
	baseTray
	caption string
}

func (self *mdTray) toString() string {
	tray := "- " + self.caption + "\n"
	for _, item := range self.tray {
		tray += item.toString() + "\n"
	}
	return tray
}

type mdPage struct {
	basePage
	title, author string
}

func (self *mdPage) output() string {
	content := "title: " + self.title + "\n"
	content += "author: " + self.author + "\n"
	for _, item := range self.content {
		content += item.toString() + "\n"
	}
	return content
}

type mdFactory struct {
}

func (self *mdFactory) createLink(caption, url string) link {
	return &mdLink{caption, url}
}
func (self *mdFactory) createTray(caption string) tray {
	return &mdTray{caption: caption}
}
func (self *mdFactory) createPage(title, author string) page {
	return &mdPage{title: title, author: author}
}
