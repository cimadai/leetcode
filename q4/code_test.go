package q4

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 2}, []int{3,4}))
	assert.Equal(t, 2, findMedianSortedArrays([]int{1, 3}, []int{2}))
}

func BenchmarkCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMedianSortedArrays([]int{1, 3}, []int{2, 4, 5})
	}
}
