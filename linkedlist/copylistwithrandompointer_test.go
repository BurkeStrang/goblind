package linkedlist

import (
	"testing"
)


func TestCopyListWithRandomPointer(t *testing.T) {
	input := [][]any{{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}
	head := buildList(input)
	copied := CopyListWithRandomPointer(head)
	if !listsEqual(head, copied) {
		t.Errorf("copied list does not match original")
	}
}
