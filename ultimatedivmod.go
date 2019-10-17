package piscine

func UltimateDivMod(x *int, y *int) {
	c := *x
	*x = *x / (*y)
	*y = c % (*y)
}
