package factory_method

type creater interface {
	createProduct(owner string) User
	registerProduct(User)
}

type User interface {
	Use() string
}

type Factory struct {
}

// インスタンス生成部分にはTemplateMethodパターンを使う。
// Goではclient-specified self patternにより親構造体のメソッド内で子構造体の
// 実装を呼ぶことができる
//
// 処理の流れが簡易でテンプレートを使わないでよければ、ファクトリとなる関数を
// 直接引数に渡してもよい
// この場合、個別のファクトリークラスを定義しなくてもよくなる。
// func (self *Factory) Create(factoryMethod func() User) {
//   self.product = factoryMethod()
//   return self.product
// }
func (self *Factory) Create(factory creater, owner string) User {
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

func (self *IDCard) Use() string {
	return self.owner
}

type IDCardFactory struct {
	*Factory
	owners []*string
}

func (self *IDCardFactory) createProduct(owner string) User {
	return &IDCard{owner}
}

func (self *IDCardFactory) registerProduct(product User) {
	owner := product.Use()
	self.owners = append(self.owners, &owner)
}
