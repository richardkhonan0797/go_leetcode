package main

import "fmt"

func main() {
	result := twoSum([]int{2, 7, 10, 13, 21}, 9)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	sisa := map[int]int{}

	for index, num := range nums {
		if val, ok := sisa[target-num]; ok {
			return []int{val, index}
		}
		sisa[num] = index
	}

	return []int{}
}
