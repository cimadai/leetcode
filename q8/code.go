package q8

import (
	"math"
	"strconv"
	"strings"
)

func isHexStr(s string) (bool, string) {
	// プレフィックス確認
	if len(s) < 3 || s[0] != '0' || s[1] != 'x' {
		return false, s[0:0]
	}
	// 含む文字列確認
	for i, c := range s[2:] {
		if c < '0' || '9' < c && c < 'A' || 'F' < c && c < 'a' || 'f' < c {
			return true, s[2:i+1]
		}
	}
	return true, s[2:]
}

func isBinaryStr(s string) (bool, string) {
	// プレフィックス確認
	if len(s) < 3 || s[0] != '0' || s[1] != 'b' {
		return false, s[0:0]
	}
	// 含む文字列確認
	for i, c := range s[2:] {
		if c < '0' || '1' < c {
			return true, s[2:i+1]
		}
	}
	return true, s[2:]
}

func isDecimalStr(s string) (bool, string) {
	skip := 0
	// プレフィックス確認
	if s[0] == '+' || s[0] == '-' {
		skip = 1
	}

	// 含む文字列確認
	for i, c := range s[skip:] {
		if c < '0' || '9' < c {
			return true, s[0:i+skip]
		}
	}
	return true, s[0:]
}

func parseInt(s string, base int) int {
	if val, err := strconv.ParseInt(s, base, 64); err != nil {
		if strings.HasSuffix(err.Error(), "value out of range") {
			if s[0] == '-' {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		} else {
			return 0
		}
	} else {
		ival := int(val)
		if math.MaxInt32 < ival {
			return math.MaxInt32
		} else if ival < math.MinInt32 {
			return math.MinInt32
		} else {
			return ival
		}
	}
}

func myAtoi(str string) int {
  	tmp := strings.TrimSpace(str)

	if len(tmp) == 0 {
		return 0
	}

	if b, s := isHexStr(tmp); b {
		return parseInt(s, 16)
	}
	if b, s := isBinaryStr(tmp); b {
		return parseInt(s, 2)
	}
	if b, s := isDecimalStr(tmp); b {
		return parseInt(s, 10)
	}

	return 0
}
