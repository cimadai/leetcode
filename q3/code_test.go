package q3

import (
	"testing"
)

func TestCode(t *testing.T) {
	ret := lengthOfLongestSubstring("abcabcbb")
	if ret != 3 {
		t.Error("ret must be 3.")
	}
}

func BenchmarkCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("abcabcbb")
	}
}