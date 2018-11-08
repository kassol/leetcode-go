package main

import (
	"fmt"
)

// BAD BAD BAD!!!!!
func lengthOfLongestSubstring(s string) int {
	for length := len(s); length > 0; length-- {
		for index := 0; index <= len(s)-length; index++ {
			if !checkRepeat(s[index : index+length]) {
				return length
			}
		}
	}
	return 0
}

func checkRepeat(s string) bool {
	m := make(map[rune]int)
	for index, value := range s {
		_, found := m[value]
		if found {
			return true
		}
		m[value] = index
	}
	return false
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring(""))
}
