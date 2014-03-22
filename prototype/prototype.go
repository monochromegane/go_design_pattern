package prototype

type producter interface {
	clone() producter
	GetName() string
}

type Manager struct {
	product producter
}

func (self *Manager) Register(producter producter) {
	self.product = producter
}

func (self *Manager) Create(name string) producter {
	producter := self.product
	return producter.clone()
}

type Product struct {
	name string
}

func (self *Product) SetUp() {
	// something takes time...
}

func (self *Product) GetName() string {
	return self.name
}

// 新しい構造体に自身の値をセットして返すことで擬似的にcloneとした。
// ポインタ参照まで考慮したdeepcopyに関しては実装もしくは機能を提供する
// パッケージが必要になる
func (self *Product) clone() producter {
	return &Product{self.name}
}
