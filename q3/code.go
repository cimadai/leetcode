package q3

func lengthOfLongestSubstring(s string) int {
	m := make([]int, 127) // ascii code mapping.
	offset := 0
	maxlen := 0

	for i, target := range s {
		x := m[target]
		if x > offset {
			offset = x
		}
		m[target] = i + 1
		if maxlen < i - offset + 1 {
			maxlen = i - offset + 1
		}
	}
	return maxlen
}
