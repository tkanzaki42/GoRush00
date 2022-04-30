package main

import (
	"piscine"
)

func main() {
	board := []string{
		"R...",
		".K..",
		"..P.",
		"....",
	}
	piscine.Checkmate(board)
	board = []string{
		"R...",
		".K..",
		"....",
		"....",
	}
	piscine.Checkmate(board)
	board = []string{
		"....",
		".K..",
		"....",
		"....",
	}
	piscine.Checkmate(board)
	board = []string{
		"....",
		".K..",
		"....",
		".R..",
	}
	piscine.Checkmate(board)
}
