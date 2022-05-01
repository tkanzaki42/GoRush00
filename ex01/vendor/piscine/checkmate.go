package piscine

func Checkmate(str []string) {
	errorMessage := Validate(str)
	if errorMessage != "" {
		PutMsg(errorMessage)
		return
	}
	x, y, size := GetKingPosition(str)
	incheck := IsInCheck(x, y, size, str)
	if incheck {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if IsMovable(x+j, y+i, size, str) {
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
