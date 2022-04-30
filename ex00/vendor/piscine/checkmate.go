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
		PutMsg("Success")
	} else {
		PutMsg("Fail")
	}
}
