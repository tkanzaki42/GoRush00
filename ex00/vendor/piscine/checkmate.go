package piscine

func Checkmate(str []string) {
	errorMessage := Validate(str)
	if errorMessage != "" {
		PutMsg(errorMessage)
		return
	}
	x, y, size := GetKingPosition(str)
	message := IsWinning(x, y, size, str)
	PutMsg(message)
}
