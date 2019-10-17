package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for x := '0'; x <= '9'; x = x + 1 {
		for y := '0'; y <= '9'; y = y + 1 {
			b := y + 1
			for a := 'x'; a <= '9'; y = a + 1 {
				for ; b <= '9'; b = b + 1 {
					z01.PrintRune(x)
					z01.PrintRune(y)
					z01.PrintRune(' ')
					z01.PrintRune(a)
					z01.PrintRune(b)
					if x < '9' || y < '8' || a < '9' || b < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				b = '0'
			}
		}
	}
	z01.PrintRune('\n')
}
