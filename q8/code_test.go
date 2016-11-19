package q8

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	assert.Equal(t, 1, myAtoi("+1"))
	assert.Equal(t, 123, myAtoi("123"))
	assert.Equal(t, 10, myAtoi("0x0A"))
	assert.Equal(t, 10, myAtoi("010"))
	assert.Equal(t, 5, myAtoi("0b101"))
	assert.Equal(t, 10, myAtoi("    010"))
	assert.Equal(t, -12, myAtoi("  -0012a42"))
	assert.Equal(t, -54, myAtoi("  -0054 23"))
	assert.Equal(t, 2147483647, myAtoi("2147483648"))
	assert.Equal(t, -2147483648, myAtoi("-2147483648"))
	assert.Equal(t, -2147483648, myAtoi("-2147483649"))
	assert.Equal(t, 2147483647, myAtoi("9223372036854775809"))
	assert.Equal(t, -2147483648, myAtoi("-9223372036854775809"))
}

//func BenchmarkCode(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		reverse(1000000003)
//	}
//}
