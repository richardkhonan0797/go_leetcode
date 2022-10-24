package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	res := ""

	for i := 0; i < len(strs[0]); i++ {
		for _, str := range strs {
			if i+1 > len(str) || string(str[i]) != string(strs[0][i]) {
				return res
			}
		}
		res = res + string(strs[0][i])
	}

	return res
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{""}))
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
}
