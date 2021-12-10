package solution1

import "sort"

func isAnagram(s, t string) bool {	
	if len(s) != len(t) {
		return false
	}
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}
