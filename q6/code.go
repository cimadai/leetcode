package q6

func convert(s string, numRows int) string {
	if (numRows == 1) {
		return s
	}

	runes := []rune(s)
	length := len(runes)
	if (length <= numRows) {
		return s
	}

	// 先頭行および末尾行の要素間の差分
	delta0 := (numRows - 1) * 2
	// 行に含まれる要素数
	lineItems := ((length + delta0 - 1) / delta0)

	idx := 0
	ret := make([]rune, length)

	for i := 0; i < numRows; i++ {
		// 残りの行数
		restRows := numRows - i - 1
		// 現在の行の次の要素までの差分
		delta1 := restRows * 2
		// 先頭行と末尾行以外は中間要素があるので、そこから列行までの差分
		delta2 := delta0 - delta1

		// 差分1の分がはみ出なければ要素数を1足す。
		items := lineItems
		if ((lineItems - 1) * delta0 + delta1) < length {
			items++
		}

		// 各行の先頭要素。
		ret[idx] = runes[i]
		idx++

		// 2個目の要素以降。
		for j := 0; j < items - 1; j++ {
			idx1 := i + delta0 * j + delta1
			idx2 := idx1 + delta2
			if delta1 > 0 && idx1 < length {
				ret[idx] = runes[idx1]
				idx++
			}
			if delta2 > 0 && idx2 < length {
				ret[idx] = runes[idx2]
				idx++
			}
		}
	}
	return string(ret)
}

