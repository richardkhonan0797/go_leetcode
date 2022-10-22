package main

import (
    "fmt"
    "math"
)

func isPalindrome(x int) bool {
    temp := x
    rev := 0
    
    for temp > 0 {
        rev = (rev * 10) + (temp % 10)
        temp /= 10
    }
    
    if rev == x {
        return true
    } else {
        return false
    }
}

func main() {
    fmt.Println(isPalindrome(121))
    fmt.Println(isPalindrome(321))
}
