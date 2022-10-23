package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int{}
	}

	i := 0
	j := len(numbers) - 1

	for i < j {
		if numbers[i]+numbers[j] > target {
			j--
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			return []int{i + 1, j + 1}
		}
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{0, -1}, -1))
}
