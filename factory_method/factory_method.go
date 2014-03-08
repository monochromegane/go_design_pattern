package factory_method

type creater interface {
	createProduct(owner string) user
	registerProduct(user)
}

type user interface {
	use() string
}

type Factory struct {
}

// インスタンス生成部分にはTemplateMethodパターンを使う。
// Goではclient-specified self patternにより親構造体のメソッド内で子構造体の
// 実装を呼ぶことができる
func (self *Factory) create(factory creater, owner string) user {
	user := factory.createProduct(owner)
	factory.registerProduct(user)
	return user
}

// 本来であれば基底クラスによるポリモフィズムを用いるが、Goの埋込による
// 構造体では機能しない。この例では親構造体に共通処理がないため、インターフェースの
// 実装によるポリモフィズムを実現する。
// よって基底クラスを定義する必要はない。
//
// type Product struct {
// }

type IDCard struct {
	owner string
}

func (self *IDCard) use() string {
	return self.owner
}

type IDCardFactory struct {
	*Factory
	owners []*string
}

func (self *IDCardFactory) createProduct(owner string) user {
	return &IDCard{owner}
}

func (self *IDCardFactory) registerProduct(product user) {
	owner := product.use()
	self.owners = append(self.owners, &owner)
}
