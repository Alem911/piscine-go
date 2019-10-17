package piscine

import "github.com/01-edu/z01"

func PrintNbr(a int) {
	if a < 0 {
		z01.PrintRune('-')
	}
	SetNbr(a)
}

func SetNbr(a int) {
	b := '0'
	if a == 0 {
		z01.PrintRune(b)
		return
	}
	for c := 1; c <= x%10; c++ {
		b++
	}
	for c := -1; c >= x%10; c-- {
		b++
	}
	if a/10 != 0 {
		SetNbr(x / 10)
	}
	z01.PrintRune(h)
	return
}
