package main

import "fmt"

var lastOccurred = make([]int, 0xffff)

func lengthOfNonRepeatingSubstr(s string) int {
	//lastOccurred := make(map[rune]int)
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubstr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubstr("bbbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubstr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubstr(""))
	fmt.Println(lengthOfNonRepeatingSubstr("b"))
	fmt.Println(lengthOfNonRepeatingSubstr("abcdefj"))
	fmt.Println(lengthOfNonRepeatingSubstr("网网王王王万"))
}