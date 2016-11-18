package q5

func getLongestMatchLength(r1 []rune, r2 []rune) int {
	match := 0
	for i1, i2 := len(r1) - 1, 0; i1 >= 0; i1, i2, match = i1-1, i2+1, match + 1 {
		if r1[i1] != r2[i2] {
			break
		}
	}
	return match
}
func getLongestMatchLength2(r []rune, halfLen int, midLen int) (int, int) {
	match := 0
	for i1, i2 := halfLen - 1, halfLen+midLen; i1 >= 0; i1, i2, match = i1-1, i2+1, match + 1 {
		if r[i1] != r[i2] {
			break
        }
	}
    if match > 0 {
        return match, midLen + match*2
    } else {
		return 0, 0
    }
}

func longestPalindrome(s string) string {
	runes := []rune(s)
	original := len(runes)
	if original == 1 {
		return s
	}

	longestLen := 0
	leng := (original * 2) - 1
	from := 0
	end := 0
	for i := 1; i < leng - 1; i++ {
		isOdd := (i & 1) == 1
        var midLen int // 中心部の長さ
		var leftLen int
		var rightLen int
		//var rightFrom int
		if isOdd {
			midLen = 0
			leftLen = (i + 1) / 2
			//rightFrom = leftLen
			rightLen = original - leftLen
		} else {
			midLen = 1
			leftLen = i / 2
			//rightFrom = leftLen + 1
			rightLen = original - leftLen - 1
		}

        // 左右の長さの短い方
		var minimumLen int
		if leftLen > rightLen {
			minimumLen = rightLen
		} else {
			minimumLen = leftLen
		}

        // 最大ここまでの長さの回文が得られそう
		probableLongest := midLen + minimumLen * 2

		if longestLen >= probableLongest {
			// すでに知ってる長さよりも短いので確かめる必要が無い
			continue
		}

        begin := leftLen - minimumLen
		r := runes[begin : begin + probableLongest]
        match, longest := getLongestMatchLength2(r, minimumLen, midLen)
		if longestLen < longest {
			// すでに知ってる長さよりも長いので改める
			longestLen = longest
			from = leftLen - match
			end = from + longest
		}
	}
	return string(runes[from:end])
}
