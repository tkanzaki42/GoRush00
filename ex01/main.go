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
	board = []string{
		"....",
		".k..",
		"....",
		".K.B",
	}
	piscine.Checkmate(board)
	board = []string{
		"....",
		".k..",
		"..N.",
		".K.B",
	}
	piscine.Checkmate(board)
}
