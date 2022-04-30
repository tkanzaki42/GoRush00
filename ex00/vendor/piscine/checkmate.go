package piscine

func Checkmate(str []string) {
	errorMessage := Validate(str)
	if errorMessage != "" {
		PutMsg(errorMessage)
		return
	}
	x, y := GetKingPosition(str)
	w, h := GetBoardSize(str)
	message := IsWinning(x, y, w, h, str)
	PutMsg(message)
}
