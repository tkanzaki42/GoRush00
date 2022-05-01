package piscine

func Validate(board []string) string {
	len_col, len_row := 0, 0
	pre_col := 0
	player_k_count := 0
	enemy_k_count := 0
	if board == nil {
		return "There is no board."
	}
	for y, row := range board {
		len_col = 0
		for _, r := range row {
			if r != 0 {
				len_col++
			}
			if r == 'K' {
				player_k_count++
			}
			if r == 'k' {
				enemy_k_count++
			}
		}
		len_row++
		if y != 0 && pre_col != len_col {
			return "Not a square."
		}
		pre_col = len_col
	}
	if len_col != len_row {
		return "Not a square."
	}
	if player_k_count != 1 {
		return "The number of player kings is wrong."
	}
	if enemy_k_count != 1 {
		return "The number of enemy kings is wrong."
	}
	return ""
}
