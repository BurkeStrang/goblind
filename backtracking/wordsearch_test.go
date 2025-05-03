package backtracking

import "testing"

func TestExist_FindsWordInBoard(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"

	if !Exist(board, word) {
		t.Errorf("Expected to find word %q in the board, but did not", word)
	}
}

func TestExist_DoesNotFindWordNotInBoard(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'B'},
	}
	word := "ABCB"

	if Exist(board, word) {
		t.Errorf("Expected not to find word %q in the board, but did", word)
	}
}

func TestExist_EmptyBoard(t *testing.T) {
	board := [][]byte{}
	word := "ANY"

	if Exist(board, word) {
		t.Errorf("Expected not to find word %q in an empty board, but did", word)
	}
}
