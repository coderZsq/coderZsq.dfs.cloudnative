package solution2

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, a := range nums {
		b := target - a
		if p, ok := hashTable[b]; ok {
			return []int{p, i}
		}
		hashTable[a] = i
	}
	return nil
}
