package solution2

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-3; i++ {
		a := nums[i]
		if i > 0 && a == nums[i-1] { // 去重 a
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			b := nums[j]
			if j > i+1 && b == nums[j-1] { // 去重 b
				continue
			}
			hashTable := map[int]struct{}{}
			for k := j + 1; k < len(nums); k++ {
				c := nums[k]
				if k > j+2 && c == nums[k-1] && nums[k-1] == nums[k-2] { // 去重 c
					continue
				}

				d := target - (a + b + c)
				if _, ok := hashTable[d]; ok {
					result = append(result, []int{a, b, c, d})
					delete(hashTable, d) // 去重 d
				} else {
					hashTable[c] = struct{}{}
				}
			}
		}
	}
	return result
}

