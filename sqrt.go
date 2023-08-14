// Write a function that returns the square root of the int passed as parameter, 
// if that square root is a whole number. Otherwise it returns 0

package piscine

func Sqrt(nb int) int {

	if nb < 0 {
		return 0	
	}

	if nb == 0 || nb == 1 {
		return nb
	}
	
	for i := 1; i*i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}