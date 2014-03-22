package flyweight

import (
	"strconv"
)

type bigChar struct {
	charName string
	fontData string
}

func NewBigChar(charName string) *bigChar {
	char := &bigChar{charName: charName}
	char.loadFontData()
	return char
}

func (self *bigChar) loadFontData() {
	num, _ := strconv.Atoi(self.charName)
	var str string
	for i := 0; i < num; i++ {
		str += "-"
	}
	self.fontData = str
}

func (self *bigChar) Print() string {
	return self.fontData
}

var instance *bigCharFactory

type bigCharFactory struct {
	pool map[string]*bigChar
}

func GetBigCharFactory() *bigCharFactory {
	if instance == nil {
		instance = &bigCharFactory{}
	}
	return instance
}

func (self *bigCharFactory) getBigChar(charName string) *bigChar {
	if self.pool == nil {
		self.pool = make(map[string]*bigChar)
	}
	if _, ok := self.pool[charName]; !ok {
		self.pool[charName] = NewBigChar(charName)
	}
	return self.pool[charName]
}

type bigString struct {
	bigChars []*bigChar
}

func NewBigString(str string) *bigString {
	bigStr := &bigString{}
	factory := GetBigCharFactory()

	for _, s := range str {
		bigStr.bigChars = append(bigStr.bigChars, factory.getBigChar(string(s)))
	}
	return bigStr
}

func (self *bigString) Print() string {
	var result string
	for _, s := range self.bigChars {
		result += s.Print() + "\n"
	}
	return result
}
