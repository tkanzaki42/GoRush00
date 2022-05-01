package main

import (
	"flag"
	"piscine"
)

func main() {
	f := flag.String("flag", "command", "command or file")
	flag.Parse()
	board := piscine.GetMap(*f)
	piscine.Checkmate(board)
	board = []string{
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
