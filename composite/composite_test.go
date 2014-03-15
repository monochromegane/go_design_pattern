package composite

import (
	"testing"
)

func TestComposite(t *testing.T) {

	rootDir := NewDirectory("root")
	usrDir := NewDirectory("usr")
	fileA := NewFile("A", 1)

	rootDir.add(usrDir)
	rootDir.add(fileA)

	fileB := NewFile("B", 2)
	usrDir.add(fileB)

	result := rootDir.printList("")

	expect := "/root (3)\n/root/usr (2)\n/root/usr/B (2)\n/root/A (1)\n"
	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}

}
