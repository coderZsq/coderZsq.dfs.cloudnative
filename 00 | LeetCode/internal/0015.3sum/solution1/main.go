package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	var result [][]int
	for i, a := range nums[:len(nums)-2] {
		if i > 0 && a == nums[i-1] { // 去重 a
			continue
		}
		hashTable := map[int]struct{}{}
		for j := i + 1; j < len(nums); j++ {
			b := nums[j]
			if j > i+2 && b == nums[j-1] && nums[j-1] == nums[j-2] { // 去重 b
				continue
			}

			c := -(a + b)
			if _, ok := hashTable[c]; ok {
				result = append(result, []int{a, b, c})
				delete(hashTable, c) // 去重 c
			} else {
				hashTable[b] = struct{}{}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{2, -2, 2, -2, 2, -2, 2, -2, 0}))
}