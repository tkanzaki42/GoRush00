package piscine

func Checkmate(str []string) {
	errorMessage := Validate(str)
	if errorMessage != "" {
		PutMsg(errorMessage)
		return
	}
	player_k_x, player_k_y, enemy_k_x, enemy_k_y, size := GetKingPosition(str)
	isEnemyInCheck := IsInCheck(enemy_k_x, enemy_k_y, size, str, true)
	isPlayerInCheck := IsInCheck(player_k_x, player_k_y, size, str, false)
	if isPlayerInCheck {
		PutMsg("You lose")
		return
	}
	if isEnemyInCheck {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if IsMovable(enemy_k_x+j, enemy_k_y+i, size, str) {
					PutMsg("Success(in check)")
					return
				}
			}
		}
		PutMsg("Success(checkmate)")
	} else {
		PutMsg("Fail")
	}
}
