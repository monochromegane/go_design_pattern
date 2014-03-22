package iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {
	bookShelf := &BookShelf{}

	asserts := []string{"A", "B", "C"}

	for _, assert := range asserts {
		bookShelf.Add(&Book{assert})
	}

	i := 0
	for iterator := bookShelf.Iterator(); iterator.HasNext(); {
		if book := iterator.Next(); book.(*Book).name != asserts[i] {
			t.Errorf("Expect Book.name to %s, but %s", asserts[i], book.(*Book).name)
		}
		i++
	}

}
