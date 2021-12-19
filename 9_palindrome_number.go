package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(isPalindrome(121))
    fmt.Println(isPalindrome(321))
}

func isPalindrome(x int) bool {
    reversed := 0
    original := x
    
    for x > 0 {
        right := x % 10
        
        power := math.Floor(math.Log10(float64(x)))
        
        reversed += right * int(math.Pow(10.0, power))
        
        x = int(math.Floor(float64(x)/10.0))
    }
    
    if reversed == original {
        return true
    }
    
    return false
}
