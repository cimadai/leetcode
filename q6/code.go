package q6

func convert(s string, numRows int) string {
	if (numRows == 1) {
		return s
	}

	length := len(s)
	if (length <= numRows) {
		return s
	}

	// 先頭行および末尾行の要素間の差分
	delta0 := (numRows - 1) * 2
	// 先頭以降のキーフレーム数(※縦に並ぶ列をキーフレームと呼ぶ)
	keyFrames := ((length + delta0 - 1) / delta0) - 1

	idx := 0
	// http://toc-dev.blogspot.jp/2012/10/go_9.html
	// 文字列を編集するときはbyteで扱って、最後のstring(bytes)するのが良い。
	// なお、元の文字列は参照だけなので[]byteや[]runeにする必要はない。
	ret := make([]byte, length, length)

	for row := 0; row < numRows; row++ {
		// 残りの行数。
		restRows := numRows - row - 1
		// 現在の行の次の要素までの差分。
		delta1 := restRows * 2
		// 先頭行と末尾行以外は中間要素があるので、そこから列行までの差分。
		delta2 := delta0 - delta1
		// 末尾行を先頭行と同じに扱いたいのでdelta1とdelta2を逆にする。
		if delta1 == 0 {
			delta1, delta2 = delta2, delta1
		}

		// 最後のおまけ(キーフレームからはみ出た部分)の有無
		extra := 0
		if row + (keyFrames * delta0) + delta1 < length {
			extra = 1
		}

		// 各行の先頭要素を並べ替え。
		ret[idx] = s[row]
		idx++

		// 2個目の要素以降を並べ替え。
		for col := 0; col < keyFrames; col++ {
			base := row + delta0 * col
			idx1 := base + delta1
			if (idx1 < length) {
				ret[idx] = s[idx1]
				idx++
			}

			// すきまを考慮。
			idx2 := idx1 + delta2
			if (delta2 > 0 && idx2 < length) {
				ret[idx] = s[idx2]
				idx++
			}
		}

		// おまけがあれば追加。
		if extra > 0 {
			base := row + delta0 * keyFrames
			idx1 := base + delta1
			ret[idx] = s[idx1]
			idx++
		}
	}
	return string(ret)
}

