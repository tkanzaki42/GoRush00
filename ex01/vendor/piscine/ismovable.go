package piscine

func IsMovable(x int, y int, size int, board []string) bool {
	if x < 0 || x > size || y < 0 || y > size {
		return false
	}
	return !IsInCheck(x, y, size, board)
}
