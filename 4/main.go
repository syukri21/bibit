package anagrams

import (
	"sort"
)

func groupAnagrams(values []string) [][]string {
	m := make(map[string][]string)
	for _, v := range values {
		bytes := []byte(v)
		sort.SliceStable(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		s := string(bytes)
		m[s] = append(m[s], v)
	}
	var ans [][]string
	for e := range m {
		ans = append(ans, m[e])
	}
	return ans
}
