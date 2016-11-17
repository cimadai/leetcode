package q5

import (
	"testing"
	"fmt"
)

func BenchmarkCode(b *testing.B) {
	fmt.Println(longestPalindrome("cbdd"))
}