package adapter

type Decorator interface {
	decorate() string
}

type Banner struct {
	str string
}

func (self *Banner) getString() string {
	return "*" + self.str + "*"
}

// 構造体の埋込による継承
type DecorateBanner struct {
	Banner
}

// 埋込による継承はgetは構造体の階層を意識しなくてもよいが
// setは明示的に埋め込んだ構造体に明示的に値を定義する必要があるため、
// 呼び出し側が階層を意識しなくてもよいようにラップするなど工夫が必要
func NewDecorateBanner(str string) *DecorateBanner {
	return &DecorateBanner{Banner{str}}
}

// インターフェースの実装と埋め込んだ構造体のメソッドによるアダプタ
func (self *DecorateBanner) decorate() string {
	return self.getString()
}
