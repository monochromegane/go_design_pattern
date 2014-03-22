package iterator

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type BookShelf struct {
	Books []*Book
}

func (self *BookShelf) Add(book *Book) {
	self.Books = append(self.Books, book)
}

// *Iterator is pointer to interface, not interface
// - Iteratorのポインタはインターフェースでないのでポインタではなくインターフェースとして戻り値を定義する
// - ポインタをレシーバーとしてインターフェースメソッドを実装しているのでIteratorの具象はポインタとして返さなければ実装したとみなされない
func (self *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{BookShelf: self}
}

type BookShelfIterator struct {
	BookShelf *BookShelf
	last      int
}

func (self *BookShelfIterator) HasNext() bool {
	if self.last >= len(self.BookShelf.Books) {
		return false
	}
	return true
}

func (self *BookShelfIterator) Next() interface{} {
	book := self.BookShelf.Books[self.last]
	self.last++
	return book
}

type Book struct {
	name string
}
