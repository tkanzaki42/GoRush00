package piscine

func GetKingPosition(input []string) (x, y int) {
	for y, _ := range input {
		for x, c := range input[y] {
			if c == 'K' {
				return x, y
			}
		}
	}
	return -1, -1
}
