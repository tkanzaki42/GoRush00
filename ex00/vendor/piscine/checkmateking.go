package piscine

func CheckmateKing(x int, y int, w int, h int, board []string) string {
	if CheckUp(x, y, w, h, board) ||
		CheckUpperRight(x, y, w, h, board) ||
		CheckRight(x, y, w, h, board) ||
		CheckLowerRight(x, y, w, h, board) ||
		CheckDown(x, y, w, h, board) ||
		CheckLowerLeft(x, y, w, h, board) ||
		CheckLeft(x, y, w, h, board) ||
		CheckUpperLeft(x, y, w, h, board) {
		return "Success"
	} else {
		return "Fail"
	}
}

func CheckLowerRight(x int, y int, w int, h int, board []string) bool {
	row := 0
	for col, _ := range board {
		if col <= y {
			continue
		}
		row = col
		if row > w {
			break
		}
		if row == x+1 && board[col][row] == 'P' {
			return true
		}
		if board[col][row] == 'Q' || board[col][row] == 'B' {
			return true
		}
	}
	return false
}

func CheckLowerLeft(x int, y int, w int, h int, board []string) bool {
	row := x
	for col, _ := range board {
		if col <= y {
			continue
		}
		row -= col - y
		if row < 0 {
			break
		}
		if row == x-1 && board[col][row] == 'P' {
			return true
		}
		if board[col][row] == 'Q' || board[col][row] == 'B' {
			return true
		}
	}
	return false
}

func CheckUpperLeft(x int, y int, w int, h int, board []string) bool {
	row := x - 1
	col := y - 1
	for {
		if col < 0 || row < 0 {
			return false
		}
		switch board[col][row] {
		case 'Q', 'B':
			return true
		case 'P', 'R':
			return false
		}
		row--
		col--
	}
	return false
}

func CheckUpperRight(x int, y int, w int, h int, board []string) bool {
	row := x + 1
	col := y - 1
	for {
		if col < 0 || row > w {
			return false
		}
		switch board[col][row] {
		case 'Q', 'B':
			return true
		case 'P', 'R':
			return false
		}
		row++
		col--
	}
}

func CheckRight(x int, y int, w int, h int, board []string) bool {
	for row, _ := range board[y] {
		if row <= x {
			continue
		}
		if board[y][row] == 'Q' || board[y][row] == 'R' {
			return true
		}
	}
	return false
}

func CheckLeft(x int, y int, w int, h int, board []string) bool {
	for row, _ := range board[y] {
		if row == x {
			break
		}
		if board[y][row] == 'Q' || board[y][row] == 'R' {
			return true
		}
	}
	return false
}

func CheckDown(x int, y int, w int, h int, board []string) bool {
	for col, _ := range board {
		if col <= y {
			continue
		}
		if board[col][x] == 'Q' || board[col][x] == 'R' {
			return true
		}
	}
	return false
}

func CheckUp(x int, y int, w int, h int, board []string) bool {
	for col, _ := range board {
		if col == y {
			break
		}
		if board[col][x] == 'Q' || board[col][x] == 'R' {
			return true
		}
	}
	return false
}
