package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {

	data, _ := ioutil.ReadFile("./p099_base_exp.txt")
	lines := strings.Split(string(data), "\n")

	var largestLineNumber int
	largestNumberSoFar := 0.0

	for i, wholeline := range lines {
		line := strings.Split(wholeline, ",")

		x, _ := strconv.ParseInt(line[0], 10, 64)
		y, _ := strconv.ParseInt(line[1], 10, 64)

		value := float64(y) * math.Log2(float64(x))
		if value > largestNumberSoFar {
			largestNumberSoFar = value
			largestLineNumber = i
		}
	}

	fmt.Printf("The line number is: %d", largestLineNumber+1)
}
