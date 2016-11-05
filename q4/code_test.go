package q4

import (
	"testing"
)

func TestCode(t *testing.T) {
	ret := findMedianSortedArrays([]int{1, 2}, []int{3,4})
	if ret != 2.5 {
		t.Errorf("ret must be 2.5. %f", ret)
	}
	ret = findMedianSortedArrays([]int{1, 3}, []int{2})
	if ret != 2 {
		t.Errorf("ret must be 2. %f", ret)
	}
}

func BenchmarkCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMedianSortedArrays([]int{1, 2}, []int{3,4})
	}
}