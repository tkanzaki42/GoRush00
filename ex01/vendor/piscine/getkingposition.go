package piscine

func GetKingPosition(input []string) (int, int, int, int, int) {
	player_k_x, player_k_y, enemy_k_x, enemy_k_y, size := -1, 0, -1, 0, 0
	for y, _ := range input {
		if player_k_x == -1 || enemy_k_x == -1 {
			for x, c := range input[y] {
				switch c {
				case 'K':
					player_k_x, player_k_y = x, y
				case 'k':
					enemy_k_x, enemy_k_y = x, y
				}
			}
		}
		size = y
	}
	return player_k_x, player_k_y, enemy_k_x, enemy_k_y, size
}
