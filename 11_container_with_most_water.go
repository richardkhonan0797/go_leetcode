package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
}

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1

	res := 0

	for start < end {
		h := math.Min(float64(height[start]), float64(height[end]))

		area := int(h) * (end - start)

		if res < area {
			res = area
		}

		if height[end] > height[start] {
			start += 1
		} else {
			end -= 1
		}
	}

	return res
}
