package main

import (
	"piscine"
)

func main() {
	board := []string{
		"....",
		".k..",
		"....",
		".KN.",
	}
	piscine.Checkmate(board)
	board = []string{
		"....",
		".k..",
		".K..",
		"....",
	}
	piscine.Checkmate(board)
}
