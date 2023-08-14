// Write a function that returns the first prime number that is equal or superior to the int passed as parameter.

// The function must be optimized in order to avoid time-outs with the tester.

// (We consider that only positive numbers can be prime numbers)

package piscine

func IsPrimeCheck(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb <= 3 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	i := 5
	for i*i <= nb {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}

	for {
		if IsPrimeCheck(nb) {
			return nb
		}
		nb++
	}
}
