package solution1

func majorityElement(nums []int) int {
	majorCnt := 0
	majorIdx := -1
	for i, num := range nums {
		cnt := count(nums, num)
		if cnt > majorCnt {
			majorCnt = cnt
			majorIdx = i
		}
	}
	return nums[majorIdx]
}

func count(nums []int, target int) int {
	cnt := 0
	for _, num := range nums {
		if num == target {
			cnt++
		}
	}
	return cnt
}