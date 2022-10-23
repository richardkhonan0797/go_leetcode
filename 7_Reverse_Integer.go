package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	res := 0

	for x != 0 {
		res = (res * 10) + (x % 10)
		x /= 10
		if res < math.MinInt32 || res > math.MaxInt32 {
			return 0
		}
	}

	return res
}

func main() {
	fmt.Println(reverse(-321))
	fmt.Println(reverse(123))
	fmt.Println(reverse(1534236469))
}
