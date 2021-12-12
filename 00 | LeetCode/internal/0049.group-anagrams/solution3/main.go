package solution3

func groupAnagrams(strs []string) [][]string {
	hashTable := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		hashTable[cnt] = append(hashTable[cnt], str)
	}
	var result [][]string
	for _, strs := range hashTable {
		result = append(result, strs)
	}
	return result
}

