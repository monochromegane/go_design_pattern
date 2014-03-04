package iterator

type Aggregator interface {
	iterator() Iterator
}

type Iterator interface {
	hasNext() bool
	next() interface{}
}

type BookShelf struct {
	Books []*Book
}

func (self *BookShelf) add(book *Book) {
	self.Books = append(self.Books, book)
}

// *Iterator is pointer to interface, not interface
// - Iteratorのポインタはインターフェースでないのでポインタではなくインターフェースとして戻り値を定義する
// - ポインタをレシーバーとしてインターフェースメソッドを実装しているのでIteratorの具象はポインタとして返さなければ実装したとみなされない
func (self *BookShelf) iterator() Iterator {
	return &BookShelfIterator{bookShelf: self}
}

type BookShelfIterator struct {
	bookShelf *BookShelf
	last      int
}

func (self *BookShelfIterator) hasNext() bool {
	if self.last >= len(self.bookShelf.Books) {
		return false
	}
	return true
}

func (self *BookShelfIterator) next() interface{} {
	book := self.bookShelf.Books[self.last]
	self.last++
	return book
}

type Book struct {
	name string
}
