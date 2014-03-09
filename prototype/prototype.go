package prototype

type producter interface {
	clone() producter
	getName() string
}

type manager struct {
	product producter
}

func (self *manager) register(producter producter) {
	self.product = producter
}

func (self *manager) create(name string) producter {
	producter := self.product
	return producter.clone()
}

type product struct {
	name string
}

func (self *product) setUp() {
	// something takes time...
}

func (self *product) getName() string {
	return self.name
}

// 新しい構造体に自身の値をセットして返すことで擬似的にcloneとした。
// ポインタ参照まで考慮したdeepcopyに関しては実装もしくは機能を提供する
// パッケージが必要になる
func (self *product) clone() producter {
	return &product{self.name}
}
