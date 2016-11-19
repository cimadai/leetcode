package q7

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	assert.Equal(t, 321, reverse(123))
	assert.Equal(t, -321, reverse(-123))
	assert.Equal(t, 1, reverse(100))
	assert.Equal(t, 0, reverse(1000000003))
	assert.Equal(t, -1, reverse(-10))
	assert.Equal(t, 0, reverse(-2147483648))

}

func BenchmarkCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse(1000000003)
	}
}
