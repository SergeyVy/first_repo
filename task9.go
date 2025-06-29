// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func reverseNumber(n int) int {
// 	isNegative := n < 0

// 	if isNegative {
// 		n = -n
// 	}

// 	str := strconv.Itoa(n)         // int → string
// 	runes := []rune(str)           // string → []rune

// 	// Разворот среза
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}

// 	reversedStr := string(runes)            // []rune → string
// 	result,_:= strconv.Atoi(reversedStr)  // string → int

// 	if isNegative {
// 		return -result
// 	}
// 	return result
// }

// func main() {
// 	fmt.Println(reverseNumber(123))   // 321
// 	fmt.Println(reverseNumber(44))    // 44
// 	fmt.Println(reverseNumber(272))   // 272
// 	fmt.Println(reverseNumber(500))   // 5
// 	fmt.Println(reverseNumber(-987))  // -789
// 	fmt.Println(reverseNumber(0))     // 0
// }