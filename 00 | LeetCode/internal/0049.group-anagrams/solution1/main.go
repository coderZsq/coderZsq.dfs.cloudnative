package main

func groupAnagrams(strs []string) [][]string {
	if len(strs) < 1 {
		return [][]string{}
	}
	var result [][]string
	result = append(result, []string{strs[0]})
	for i := 1; i< len(strs); i++ {
		str := strs[i]
		addAnagram := false
		for idx, strs := range result {
			if isAnagram(str, strs[0]) {
				strs = append(strs, str)
				result[idx] = strs
				addAnagram = true
				break
			}
		}
		if !addAnagram {
			result = append(result, []string{str})
		}
	}
	return result
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}
