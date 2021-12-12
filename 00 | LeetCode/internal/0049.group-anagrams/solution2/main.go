package solution2

import "sort"

func groupAnagrams(strs []string) [][]string {
	hashTable := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		hashTable[sortedStr] = append(hashTable[sortedStr], str)
	}
	var result [][]string
	for _, strs := range hashTable {
		result = append(result, strs)
	}
	return result
}
