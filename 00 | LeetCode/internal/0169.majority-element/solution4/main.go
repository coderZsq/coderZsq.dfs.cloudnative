package solution4

func majorityElement(nums []int) int {
	return majorityElementRec(nums, 0, len(nums)-1)
}

func majorityElementRec(nums []int, l int, r int) int {
	if l == r {
		return nums[l]
	}

	mid := (r-l)>>1 + l
	left := majorityElementRec(nums, l, mid)
	right := majorityElementRec(nums, mid+1, r)

	if left == right {
		return left
	}

	leftCnt := countInRange(nums, left, l, r)
	rightCnt := countInRange(nums, right, l, r)

	if leftCnt > rightCnt {
		return left
	} else {
		return right
	}
}

func countInRange(nums []int, target int, l int, r int) int {
	cnt := 0
	for i := l; i <= r; i++ {
		if nums[i] == target {
			cnt++
		}
	}
	return cnt
}
