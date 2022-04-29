package piscine

import "ft"

func PutError(err string) {
	for _, s := range err {
		ft.PrintRune(s)
	}
	ft.PrintRune('\n')
}
