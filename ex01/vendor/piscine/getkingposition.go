package piscine

func GetKingPosition(input []string) (x, y, size int) {
	king_x, king_y, size := -1, 0, 0
	for y, _ := range input {
		if king_x == -1 {
			for x, c := range input[y] {
				if c == 'K' {
					king_x, king_y = x, y
				}
			}
		}
		size = y
	}
	return king_x, king_y, size
}
