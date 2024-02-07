package main

import (
	"fmt"

	"github.com/mecozma/go-testing/math"
)

func main() {
	sum := math.Sum([]int{10, -2, 3})
	if sum != 11 {
		msg := fmt.Sprintf("FAIL: Wanted 11, but got %d\n", sum)
		panic(msg)
	}
	add := math.Add(5, 10)
	if add != 15 {
		msg := fmt.Sprintf("FAIL: Wanted 15 but received %d/n", add)
		panic(msg)
	}
	fmt.Println("PASS")
}
