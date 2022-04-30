package piscine

func Validate(board []string) string{
	len_col, len_row := 0, 0
	pre_col := 0
	k_count := 0
	for y, row := range board {
		len_col = 0
		for _, r := range row {
			if r != 0 {
				len_col++
			}
			if r == 'K' {
				k_count++
			}
		}
		len_row++
		if y != 0 && pre_col != len_col {
			return "Not a square."
		}
		pre_col = len_col
	}
	if len_col != len_row {
		return "Not a square"
	}
	if k_count != 1 {
		return "The number of kings is wrong."
	}
	return ""
}
