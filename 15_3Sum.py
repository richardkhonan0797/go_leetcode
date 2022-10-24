package main

import (
	"fmt"
)

func bubbleSort(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}

func threeSum(nums []int) [][]int {
	res := [][]int{}

	if len(nums) < 3 {
		return res
	}

	sorted := bubbleSort(nums)

	for i := 0; i < len(sorted)-2; i++ {
		if i > 0 && sorted[i] == sorted[i-1] {
			continue
		}
		j := i + 1
		k := len(sorted) - 1

		for j < k {
			if sorted[i]+sorted[j]+sorted[k] > 0 {
				k--
			} else if sorted[i]+sorted[j]+sorted[k] < 0 {
				j++
			} else {
				res = append(res, [][]int{{sorted[i], sorted[j], sorted[k]}}...)
				j++
				for sorted[j] == sorted[j-1] && j < k {
					j++
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
