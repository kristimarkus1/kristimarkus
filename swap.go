package piscine

func Swap(a *int, b *int) {
	kohatäitja := *a
	*a = *b
	*b = kohatäitja
}
