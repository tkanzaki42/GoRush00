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
	col := y + 1
	row := x + 1
	for col <= h {
		if row > w {
			break
		}
		if row == x+1 && board[col][row] == 'P' {
			return true
		}
		switch board[col][row] {
		case 'R':
			return false
		case 'P':
			return false
		case 'Q':
			return true
		case 'B':
			return true
		}
		row++
		col++
	}
	return false
}

func CheckLowerLeft(x int, y int, w int, h int, board []string) bool {
	col := y + 1
	row := x - 1
	for col <= h {
		if row < 0 {
			break
		}
		if row == x-1 && board[col][row] == 'P' {
			return true
		}
		switch board[col][row] {
		case 'R':
			return false
		case 'P':
			return false
		case 'Q':
			return true
		case 'B':
			return true
		}
		row--
		col++
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
	col := y
	row := x + 1
	for row <= w {
		switch board[col][row] {
		case 'R':
			return true
		case 'P':
			return false
		case 'Q':
			return true
		case 'B':
			return false
		}
		row++
	}
	return false
}

func CheckLeft(x int, y int, w int, h int, board []string) bool {
	col := y
	row := x - 1
	for row >= 0 {
		switch board[col][row] {
		case 'R':
			return true
		case 'P':
			return false
		case 'Q':
			return true
		case 'B':
			return false
		}
		row--
	}
	return false
}

func CheckDown(x int, y int, w int, h int, board []string) bool {
	col := y + 1
	row := x
	for col <= h {
		switch board[col][row] {
		case 'R':
			return true
		case 'P':
			return false
		case 'Q':
			return true
		case 'B':
			return false
		}
		col++
	}
	return false
}

func CheckUp(x int, y int, w int, h int, board []string) bool {
	col := y - 1
	row := x
	for col >= 0 {
		switch board[col][row] {
		case 'R':
			return true
		case 'P':
			return false
		case 'Q':
			return true
		case 'B':
			return false
		}
		col--
	}
	return false
}
