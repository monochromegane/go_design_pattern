package adapter

type DelegateDecorateBanner struct {
	// 埋込ではなく構造体のフィールドとしてadapteeを保持
	banner *Banner
}

func NewDelegateDecorateBanner(str string) *DelegateDecorateBanner {
	return &DelegateDecorateBanner{&Banner{str}}
}

// インターフェースの実装と移譲によるアダプタ
func (self *DelegateDecorateBanner) decorate() string {
	return self.banner.getString()
}
