package q6

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	//assert.Equal(t, "0246813579", convert("0123456789", 2))
	//assert.Equal(t, "0481357926", convert("0123456789", 3))
	//assert.Equal(t, "0615724839", convert("0123456789", 4))
	//assert.Equal(t, "0817926354", convert("0123456789", 5))
	assert.Equal(t, "06157b248a39", convert("0123456789ab", 4))
	//assert.Equal(t, "A", convert("A", 1))
	//assert.Equal(t, "AB", convert("AB", 1))
}

func BenchmarkCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convert("0123456789ab", 4)
	}
}
