package backtracking

func Exist(board [][]byte, word string) bool {

	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	rows := len(board)
	cols := len(board[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for i := range rows {
		for j := range cols {
			if board[i][j] == word[0] && search(i, j, 0, word, board, visited) {
				return true
			}
		}
	}

	return false
}

func search(r, c, index int, word string, board [][]byte, visited [][]bool) bool {
	rows := len(board)
	cols := len(board[0])

	if index == len(word) {
		return true
	}
	if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != word[index] || visited[r][c] {
		return false
	}

	visited[r][c] = true
	result := search(r+1, c, index+1, word, board, visited) ||
		search(r-1, c, index+1, word, board, visited) ||
		search(r, c+1, index+1, word, board, visited) ||
		search(r, c-1, index+1, word, board, visited)

	visited[r][c] = false
	return result
}
