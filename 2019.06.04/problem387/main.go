package main

import (
	"fmt"
	"math"
	"math/big"
)

var maxLimit = int64(math.Pow10(14))

var isRightTruncatableHarshadNumCache = map[int64]bool{}
var maxCached int64
var primeMap = map[int64]bool{}
var maxPrime int64

func sumOfDigits(n int64) (sum int64) {
	var d int64
	for n > 0 {
		d = n % 10
		n = n / 10
		sum += d
	}
	return
}

func isHarshadNum(n int64) bool {
	//if isRightTruncatableHarshadNumCache[n] == true {
	//return true
	//}

	//if n < maxCached {
	//return false
	//}

	is := n != 0 && (n < 10 || n%sumOfDigits(n) == 0)
	//if is {
	//isRightTruncatableHarshadNumCache[n] = true
	//maxCached = n
	//}
	return is
}

func isRightTruncatableHarshadNum(n int64) bool {
	if n == 0 {
		return false
	}
	if n < 10 {
		return true
	}

	return isHarshadNum(n) && isRightTruncatableHarshadNum(n/10)
}

func isStrongRightTruncatableHN(n int64) bool {
	return isRightTruncatableHarshadNum(n) && isPrime(n/sumOfDigits(n))
}

func isPrime(n int64) bool {
	return big.NewInt(n).ProbablyPrime(3)
}

func isStrongHarshadNumber(n int64) bool {
	return isHarshadNum(n) && isPrime(n/sumOfDigits(n))
}

func recursiveBuilder(n int64) int64 {
	//fmt.Printf("something: %d\n", n)
	if n > maxLimit {
		return 0
	}
	var i, sum int64
	n10 := n * 10
	for i = 0; i <= 9; i++ {
		j := n10 + i
		if isRightTruncatableHarshadNum(j) {
			sum += recursiveBuilder(j)
		} else {
			if j <= maxLimit && isPrime(j) && isStrongHarshadNumber(n) {
				fmt.Printf("j: %d\n", j)
				sum += j
			}
		}
	}
	return sum
}

func sumOfSRTHN(limit int64) (sum int64) {
	var number int64
	fmt.Println(limit)
	for number = 3; number < limit; number += 1 {
		if number%3 == 0 || number%5 == 0 {
			continue
		}
		if isStrongRightTruncatableHN(number/10) && isPrime(number) {
			fmt.Printf("%d\n", number)
			sum += number
		}
	}

	return
}

func main() {
	//fmt.Printf("90619 ==? %d\n", sumOfSRTHN(int64(math.Pow10(4))))
	var sum int64
	for i := 1; i <= 9; i++ {
		s := recursiveBuilder(int64(i))
		sum += int64(s)
	}
	fmt.Printf("Sum is: %d\n", sum)
}
