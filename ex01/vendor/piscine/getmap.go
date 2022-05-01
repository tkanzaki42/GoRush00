package piscine

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
)

func GetMap(fl string) []string {
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
