package solution2

func majorityElement(nums []int) int {
	hashTable := map[int]int{}
	for _, num := range nums {
		hashTable[num]++
	}
	majorCnt := 0
	majorNum := 0
	for k, v := range hashTable {
		if majorCnt < v {
			majorCnt = v
			majorNum = k
		}
	}
	return majorNum
}
