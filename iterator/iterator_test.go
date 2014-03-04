package iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {
	bookShelf := &BookShelf{}

	asserts := []string{"A", "B", "C"}

	for _, assert := range asserts {
		bookShelf.add(&Book{assert})
	}

	i := 0
	for iterator := bookShelf.iterator(); iterator.hasNext(); {
		if book := iterator.next(); book.(*Book).name != asserts[i] {
			t.Errorf("Expect Book.name to %s, but %s", asserts[i], book.(*Book).name)
		}
		i++
	}

}
