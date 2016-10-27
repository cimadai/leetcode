package q1

import (
	"testing"
)

func TestCode(t *testing.T) {
	ret := twoSum([]int{3, 2, 4}, 6)
	if len(ret) != 2 {
		t.FailNow()
	}

	if ret[0] != 1 {
		t.FailNow()
	}

	if ret[1] != 2 {
		t.FailNow()
	}
}

func BenchmarkCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{3, 2, 4}, 6)
	}
}