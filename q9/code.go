package q9

import (
	"strconv"
)

func reverse(x int) int {
	bytes := []byte(strconv.Itoa(x))
	length := len(bytes)

	from := 0
	med := length / 2

	// 符号を除く、両端から文字列をひっくり返す
	for i, j := from, length - 1; i < med; i, j = i + 1, j - 1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	// パースできなければ0を、パースできていたらInt32まではそれを、オーバーフローしていたら0を返す。
	if val, err := strconv.Atoi(string(bytes)); err != nil {
		return 0
	} else {
		MaxInt32 := 1<<31 - 1
		if val < -MaxInt32 || MaxInt32 < val {
			return 0
		} else {
			return val
		}
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false;
	}

	return x == reverse(x)
}
