package piscine

func GetBoardSize(input []string) (w, h int) {
	count_w := 0
	for w := range input {
		count_w = w
	}
	count_h := 0
	for h := range input[0] {
		count_h = h
	}
	return count_w, count_h
}
