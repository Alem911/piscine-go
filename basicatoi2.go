package piscine

func BasicAtoi2(s string) int {
	runes := []rune(s)
	num := 0
	length := 0
	for i := range runes {
		length = i
	}
	ten := 1
	for j := 0; j < length; j++ {
		ten *= 10
	}
	for i := range runes {
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		temp := 0
		for k := '0'; k < runes[i]; k++ {
			temp++
		}
		num += temp * ten
		ten /= 10
	}
	return num
}
