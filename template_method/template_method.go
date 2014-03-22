package template_method

type printer interface {
	open() string
	print() string
	close() string
}

type AbstractDisplay struct {
}

// 構造体をレシーバとするメソッドの引数に同じ構造体を渡す必要がある。
//
// なぜか？
// - Go言語の構造体埋込と継承は別物
// - 子構造体を親構造体型の変数に格納することはできない（そもそも親子関係ではなく構造体を内包しているだけ）
// - 親構造体のメソッド内で子構造体を呼び戻すことはしない
//  type Abstract struct { some_interface }
//  func (self *Abstract) templateMethod() {
//      self.method() // panic: runtime error: invalid memory address or nil pointer dereference
//  }
//
// どうするか？
// 構造体をレシーバとするメソッドの引数に同じ構造体を渡すパターン（client-specified self pattern）を使う
// これにより親構造体のメソッド内で子構造体の実装を呼ぶことができる

func (self *AbstractDisplay) Display(printer printer) string {
	result := printer.open()
	for i := 0; i < 5; i++ {
		result += printer.print()
	}
	result += printer.close()
	return result
}

type CharDisplay struct {
	*AbstractDisplay
	Char rune
}

func (self *CharDisplay) open() string {
	return "<<"
}
func (self *CharDisplay) print() string {
	return string(self.Char)
}
func (self *CharDisplay) close() string {
	return ">>"
}

type StringDisplay struct {
	*AbstractDisplay
	Str string
}

func (self *StringDisplay) open() string {
	return self.printLine()
}
func (self *StringDisplay) print() string {
	return "| " + self.Str + " |\n"
}
func (self *StringDisplay) close() string {
	return self.printLine()
}

func (self *StringDisplay) printLine() string {
	line := "+-"
	for _, _ = range self.Str {
		line += "-"
	}
	line += "-+\n"
	return line
}

// 継承を利用しないTemplateMethodパターン

// Goでは継承を利用したパターンは実装がいびつになるのでオススメしない。
// 例えばテンプレートメソッドをもつ構造体とメソッドを持つインターフェースを
// 分けて定義して処理を移譲する構成にすればスッキリする
// が、これは継承を使っていないのでTemplateMethodパターンというよりは
// Strategyパターンに近い

// type Display struct {
//         printer printer
// }
//
// func (self *Display) Display() string {
// 	result := self.printer.open()
// 	for i := 0; i < 5; i++ {
// 		result += self.printer.print()
// 	}
// 	result += self.printer.close()
// 	return result
// }
