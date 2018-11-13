package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	m := make([]int, 128)
	max := 0
	length := len(s)
	for i, j := 0, 0; j < length; j++ {
		if m[s[j]] > i {
			i = m[s[j]]
		}
		if j-i+1 > max {
			max = j - i + 1
		}
		m[s[j]] = j + 1
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring(""))
}
