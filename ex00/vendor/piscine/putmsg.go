package piscine

import "ft"

func PutMsg(err string) {
	for _, s := range err {
		ft.PrintRune(s)
	}
	ft.PrintRune('\n')
}
