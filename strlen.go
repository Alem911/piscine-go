package piscine

func StrLen(str string) int {

	w := 0
	for _, i := range str {
		if i == i {
			w++
		}
	}
	return w
}
