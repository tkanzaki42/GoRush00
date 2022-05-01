package main

import (
	"flag"
	"io/ioutil"
	"os"
	"piscine"
	"strings"
)

func main() {
	f := flag.String("flag", "command", "command or file")
	flag.Parse()
	board := getMap(*f)
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

func getMap(fl string) []string {
	switch fl {
	case "command":
		if len(flag.Args()) == 0 {
			return nil
		}
		input := make([]string, len(flag.Args()))
		for i := 0; i < len(flag.Args()); i++ {
			input[i] = strings.Clone(flag.Args()[i])
		}
		return input
	case "file":
		if len(flag.Args()) == 0 {
			return nil
		}
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			return nil
		}
		defer f.Close()
		input, err := ioutil.ReadAll(f)
		if err != nil {
			return nil
		}
		if string(input) == "" {
			return nil
		}
		board := strings.Split(string(input), "\n")
		for i := range board {
			board[i] = strings.ReplaceAll(board[i], "\r", "")
		}
		return board
	default:
		return nil
	}
}
