package main

import (
	"fmt"
	"math"
	"math/bits"
)

func Factorial(n int) (result float64) {
	if n > 0 {
		result = float64(n) * Factorial(n-1)
		return result
	}
	return 1.0
}

func twosPow(n int) int {
	return int(math.Pow(2, float64(n)))
}

func problem(n int) float64 {
	upto := twosPow(n) + 1

	f := Factorial(n + 1)
	p := 0.0
	for perm := 0; perm < upto; perm++ {
		wcount := bits.OnesCount(uint(perm))
		if wcount <= n/2 {
			continue
		}

		permProb := f

		for b := 0; b < n; b++ {
			sequence := 1 << uint(b)
			if perm&sequence == 0 {
				permProb *= float64(b+1) / float64(b+2)
			} else {
				permProb *= 1.0 / float64(b+2)
			}
		}
		p += permProb
	}

	return p
}

func main() {
	fmt.Printf("The case for 4 is %f\n", problem(4))
	fmt.Printf("The case for 10 is %f\n", problem(10))
	fmt.Printf("The case for 15 is %f\n", problem(15))
}
