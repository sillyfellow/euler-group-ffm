package main

import (
	"fmt"
	"math/big"
)

func sumOfDigits(n *big.Int) (sum int64) {

	zero := big.NewInt(0)
	ten := big.NewInt(10)
	d := big.NewInt(0)

	for n.Cmp(zero) == 1 {
		n, d = n.QuoRem(n, ten, d)
		sum += d.Int64()
	}

	return
}

func main() {
	var result big.Int

	exp := big.NewInt(1000)
	base := big.NewInt(2)

	fmt.Printf("%v\n", result.Exp(base, exp, nil))
	fmt.Printf("%v\n", sumOfDigits(result.Exp(base, exp, nil)))
}
