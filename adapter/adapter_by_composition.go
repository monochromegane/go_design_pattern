package adapter

type ComposiotionDecorateBanner struct {
	// 埋込ではなく構造体のフィールドとしてadapteeを保持
	banner *Banner
}

func NewCompositionDecorateBanner(str string) *ComposiotionDecorateBanner {
	return &ComposiotionDecorateBanner{&Banner{str}}
}

// インターフェースの実装と移譲によるアダプタ
func (self *ComposiotionDecorateBanner) Decorate() string {
	return self.banner.getString()
}
