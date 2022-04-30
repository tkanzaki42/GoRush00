package piscine

func Checkmate(str []string) {
	errorMessage := Validate(str)
	if errorMessage != "" {
		PutError(errorMessage)
		return
	}
	x, y := GetKingPosition(str)
	w, h := GetBoardSize(str)
	message := CheckmateKing(x, y, w, h, str)
	PutError(message)
}
