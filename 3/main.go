package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	length := len(s)
	ann, i, j := 0, 0, 0
	for i < length && j < length {
		index, found := m[s[j]]
		if found && index >= i {
			i = index + 1
		} else {
			m[s[j]] = j
			j++
			if j-i > ann {
				ann = j - i
			}
		}
	}
	return ann
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring(""))
}
