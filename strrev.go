package piscine

func StrRev(s string) string {
	runes := []rune(s)
	length := 0
	for i := range runes {
		length = i
	}

	for i, j := 0, length; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
