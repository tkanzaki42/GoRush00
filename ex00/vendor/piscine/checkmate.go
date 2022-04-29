package piscine

func Checkmate(str []string) {
	errorMessage := Validate(str)
	if errorMessage != "" {
		PutError(errorMessage)
		return
	}
	x, y := GetKingPosition(str)
	CheckmateKing(x, y, str)
}
