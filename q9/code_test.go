package q9

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	assert.Equal(t, true, isPalindrome(101))
	assert.Equal(t, true, isPalindrome(0))
}

func BenchmarkCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome(123454321)
	}
}
