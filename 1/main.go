package main

import (
	"log"
	"time"
)

func twoSum(nums []int, target int) []int {
	for i, value := range nums {
		for j, value2 := range nums[i+1:] {
			if value+value2 == target {
				return []int{i, i + j + 1}
			}
		}
	}
	return []int{}
}

func twoSumBest(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		t, found := m[target-num]
		if found {
			return []int{t, i}
		}
		m[num] = i
	}
	return nil
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
	timeTrack(twoSum, []int{2, 7, 11, 15}, 26, "mine")
	timeTrack(twoSumBest, []int{2, 7, 11, 15}, 26, "best")
}
