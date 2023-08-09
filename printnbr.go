package piscine

import (
	"math/big"

	"github.com/01-edu/z01"
)

func PrintNbr(n int64) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	bigIntN := new(big.Int).SetInt64(n)
	if n < 0 {
		z01.PrintRune('-')
		bigIntN.Abs(bigIntN)
	}

	ten := big.NewInt(10)
	zero := big.NewInt(0)
	reverse := new(big.Int)

	for bigIntN.Cmp(zero) > 0 {
		remainder := new(big.Int)
		bigIntN, remainder = bigIntN.DivMod(bigIntN, ten, remainder)
		reverse = reverse.Mul(reverse, ten)
		reverse = reverse.Add(reverse, remainder)
	}

	for reverse.Cmp(zero) > 0 {
		remainder := new(big.Int)
		reverse, remainder = reverse.DivMod(reverse, ten, remainder)
		z01.PrintRune(rune(remainder.Int64() + '0'))
	}
}
