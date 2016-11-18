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
		var leftLen int
		var rightLen int
		var rightFrom int
		if isOdd {
			leftLen = (i + 1) / 2
			rightFrom = leftLen
			rightLen = original - leftLen
		} else {
			leftLen = i / 2
			rightFrom = leftLen + 1
			rightLen = original - leftLen - 1
		}
		var minimumLen int
		if leftLen > rightLen {
			minimumLen = rightLen
		} else {
			minimumLen = leftLen
		}
		probableLongest := rightFrom - leftLen + minimumLen * 2

		if longestLen >= probableLongest {
			// すでに知ってる長さよりも短いので確かめる必要が無い
			continue
		}

		//fmt.Printf("i = %d, leftLen = %d, rightLen = %d, rightFrom = %d, minimumLen = %d\n", i, leftLen, rightLen, rightFrom, minimumLen)
		r1 := runes[leftLen - minimumLen : leftLen]
		r2 := runes[rightFrom : rightFrom + minimumLen]
		longestMatch := getLongestMatchLength(r1, r2)
		f := leftLen - longestMatch
		e := rightFrom + longestMatch
		longest := e - f

		if longestLen < longest {
			// すでに知ってる長さよりも長いので改める
			longestLen = longest
			from = f
			end = e
		}
		//fmt.Printf("longestLen = %d, r1 = %s, r2 = %s\n", longestLen, string(r1), string(r2))
	}
	return string(runes[from:end])
}
