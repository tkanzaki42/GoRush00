package piscine

func IsInCheck(x int, y int, size int, board []string, enemyMode bool) bool {
	w, h := size, size
	if CheckKing(x, y, w, h, board, enemyMode) ||
		CheckKnight(x, y, w, h, board, enemyMode) ||
		CheckUp(x, y, w, h, board, enemyMode) ||
		CheckUpperRight(x, y, w, h, board, enemyMode) ||
		CheckRight(x, y, w, h, board, enemyMode) ||
		CheckLowerRight(x, y, w, h, board, enemyMode) ||
		CheckDown(x, y, w, h, board, enemyMode) ||
		CheckLowerLeft(x, y, w, h, board, enemyMode) ||
		CheckLeft(x, y, w, h, board, enemyMode) ||
		CheckUpperLeft(x, y, w, h, board, enemyMode) {
		return true
	} else {
		return false
	}
}

func CheckKing(x int, y int, w int, h int, board []string, enemyMode bool) bool {
	k := ' '
	if enemyMode {
		k = 'K'
	} else {
		k = 'k'
	}
	if x <= w && y-1 >= 0 && rune(board[y-1][x]) == k {
		return true
	} else if x+1 <= w && y-1 >= 0 && rune(board[y-1][x+1]) == k {
		return true
	} else if x+1 <= w && y <= h && rune(board[y][x+1]) == k {
		return true
	} else if x+1 <= w && y+1 <= h && rune(board[y+1][x+1]) == k {
		return true
	} else if x >= 0 && y+1 <= h && rune(board[y+1][x]) == k {
		return true
	} else if x-1 >= 0 && y+1 <= h && rune(board[y+1][x-1]) == k {
		return true
	} else if x-1 >= 0 && y >= 0 && rune(board[y][x-1]) == k {
		return true
	} else if x-1 >= 0 && y-1 >= 0 && rune(board[y-1][x-1]) == k {
		return true
	} else {
		return false
	}
}

func CheckKnight(x int, y int, w int, h int, board []string, enemyMode bool) bool {
	n := ' '
	if enemyMode {
		n = 'N'
	} else {
		n = 'n'
	}
	if x+1 <= w && y-2 >= 0 && rune(board[y-2][x+1]) == n {
		return true
	} else if x+2 <= w && y-1 >= 0 && rune(board[y-1][x+2]) == n {
		return true
	} else if x+2 <= w && y+1 <= h && rune(board[y+1][x+2]) == n {
		return true
	} else if x+1 <= w && y+2 <= h && rune(board[y+2][x+1]) == n {
		return true
	} else if x-1 >= 0 && y+2 <= h && rune(board[y+2][x-1]) == n {
		return true
	} else if x-2 >= 0 && y+1 <= h && rune(board[y+1][x-2]) == n {
		return true
	} else if x-2 >= 0 && y-1 >= 0 && rune(board[y-1][x-2]) == n {
		return true
	} else if x-1 >= 0 && y-2 >= 0 && rune(board[y-2][x-1]) == n {
		return true
	} else {
		return false
	}
}
func CheckLowerRight(x int, y int, w int, h int, board []string, enemyMode bool) bool {
	r, p, q, b := ' ', ' ', ' ', ' '
	if enemyMode {
		r = 'R'
		p = 'P'
		q = 'Q'
		b = 'B'
	} else {
		r = 'r'
		p = 'p'
		q = 'q'
		b = 'b'
	}
	col := y + 1
	row := x + 1
	for col <= h {
		if row > w {
			break
		}
		if row == x+1 && rune(board[col][row]) == p {
			return true
		}
		switch rune(board[col][row]) {
		case r:
			return false
		case p:
			return false
		case q:
			return true
		case b:
			return true
		}
		row++
		col++
	}
	return false
}

func CheckLowerLeft(x int, y int, w int, h int, board []string, enemyMode bool) bool {
	r, p, q, b := ' ', ' ', ' ', ' '
	if enemyMode {
		r = 'R'
		p = 'P'
		q = 'Q'
		b = 'B'
	} else {
		r = 'r'
		p = 'p'
		q = 'q'
		b = 'b'
	}
	col := y + 1
	row := x - 1
	for col <= h {
		if row < 0 {
			break
		}
		if row == x-1 && rune(board[col][row]) == p {
			return true
		}
		switch rune(board[col][row]) {
		case r:
			return false
		case p:
			return false
		case q:
			return true
		case b:
			return true
		}
		row--
		col++
	}
	return false
}

func CheckUpperLeft(x int, y int, w int, h int, board []string, enemyMode bool) bool {
	r, p, q, b := ' ', ' ', ' ', ' '
	if enemyMode {
		r = 'R'
		p = 'P'
		q = 'Q'
		b = 'B'
	} else {
		r = 'r'
		p = 'p'
		q = 'q'
		b = 'b'
	}
	row := x - 1
	col := y - 1
	for {
		if col < 0 || row < 0 {
			return false
		}
		switch rune(board[col][row]) {
		case r:
			return false
		case p:
			return false
		case q:
			return true
		case b:
			return true
		}
		row--
		col--
	}
	return false
}

func CheckUpperRight(x int, y int, w int, h int, board []string, enemyMode bool) bool {
	r, p, q, b := ' ', ' ', ' ', ' '
	if enemyMode {
		r = 'R'
		p = 'P'
		q = 'Q'
		b = 'B'
	} else {
		r = 'r'
		p = 'p'
		q = 'q'
		b = 'b'
	}
	row := x + 1
	col := y - 1
	for {
		if col < 0 || row > w {
			return false
		}
		switch rune(board[col][row]) {
		case r:
			return false
		case p:
			return false
		case q:
			return true
		case b:
			return true
		}
		row++
		col--
	}
}

func CheckRight(x int, y int, w int, h int, board []string, enemyMode bool) bool {
	r, p, q, b := ' ', ' ', ' ', ' '
	if enemyMode {
		r = 'R'
		p = 'P'
		q = 'Q'
		b = 'B'
	} else {
		r = 'r'
		p = 'p'
		q = 'q'
		b = 'b'
	}
	col := y
	row := x + 1
	for row <= w {
		switch rune(board[col][row]) {
		case r:
			return true
		case p:
			return false
		case q:
			return true
		case b:
			return false
		}
		row++
	}
	return false
}

func CheckLeft(x int, y int, w int, h int, board []string, enemyMode bool) bool {
	r, p, q, b := ' ', ' ', ' ', ' '
	if enemyMode {
		r = 'R'
		p = 'P'
		q = 'Q'
		b = 'B'
	} else {
		r = 'r'
		p = 'p'
		q = 'q'
		b = 'b'
	}
	col := y
	row := x - 1
	for row >= 0 {
		switch rune(board[col][row]) {
		case r:
			return true
		case p:
			return false
		case q:
			return true
		case b:
			return false
		}
		row--
	}
	return false
}

func CheckDown(x int, y int, w int, h int, board []string, enemyMode bool) bool {
	r, p, q, b := ' ', ' ', ' ', ' '
	if enemyMode {
		r = 'R'
		p = 'P'
		q = 'Q'
		b = 'B'
	} else {
		r = 'r'
		p = 'p'
		q = 'q'
		b = 'b'
	}
	col := y + 1
	row := x
	for col <= h {
		switch rune(board[col][row]) {
		case r:
			return true
		case p:
			return false
		case q:
			return true
		case b:
			return false
		}
		col++
	}
	return false
}

func CheckUp(x int, y int, w int, h int, board []string, enemyMode bool) bool {
	r, p, q, b := ' ', ' ', ' ', ' '
	if enemyMode {
		r = 'R'
		p = 'P'
		q = 'Q'
		b = 'B'
	} else {
		r = 'r'
		p = 'p'
		q = 'q'
		b = 'b'
	}
	col := y - 1
	row := x
	for col >= 0 {
		switch rune(board[col][row]) {
		case r:
			return true
		case p:
			return false
		case q:
			return true
		case b:
			return false
		}
		col--
	}
	return false
}
