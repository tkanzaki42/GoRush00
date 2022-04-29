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
}
