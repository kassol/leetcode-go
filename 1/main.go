package main

import (
	"log"
	"time"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		j, found := m[target-v]
		if found {
			return []int{j, i}
		}
		m[v] = i
	}
	return []int{}
}

type fn func(num []int, target int) []int

func timeTrack(f fn, num []int, target int, name string) {
	start := time.Now()
	result := f(num, target)
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
	log.Println(result)
}

func main() {
	timeTrack(twoSum, []int{2, 7, 11, 15}, 26, "new")
}
