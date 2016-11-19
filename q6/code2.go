package q6

import (
	"math"
	"sync"
)
func convert2(s string, numRows int) string {
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

	//idx := 0
	ret := make([]rune, length)
	next := 0

	wg := &sync.WaitGroup{}  // WaitGroupの値を作る
	for i := 0; i < numRows; i++ {
		// 残りの行数
		restRows := numRows - i - 1
		// 現在の行の次の要素までの差分
		delta1 := restRows * 2
		// 先頭行と末尾行以外は中間要素があるので、そこから列行までの差分
		delta2 := delta0 - delta1
		if delta1 == 0 {
			delta1, delta2 = delta2, delta1
		}

		// キーフレーム分
		keyFrames := int(math.Ceil((float64(length - i - 1) + 0.999999) / float64(delta0)))
		// すき間分
		intermediates := keyFrames - 1
		if i == 0 || i == numRows - 1 {
			intermediates = 0
		}
		// 最後のおまけ
		extra := 0
		if i + ((keyFrames - 1) * delta0) + delta1 < length {
			extra = 1
		}

		wg.Add(1)  // wgをインクリメント
		go convertImpl(wg, runes, ret, i, keyFrames, intermediates, extra, next, delta0, delta1, delta2)
		next += keyFrames + intermediates + extra
	}
	wg.Wait()
	return string(ret)
}

func convertImpl(wg *sync.WaitGroup, runes []rune, output []rune, line int, keyFrames int, intermediates int, extra int, fromIdx int, delta0 int, delta1 int, delta2 int) {
	// 各行の先頭要素。
	output[fromIdx] = runes[line]
	fromIdx++

	if intermediates == 0 {
		// 2個目の要素以降。
		for j := 0; j < keyFrames - 1; j++ {
			base := line + delta0 * j
			idx1 := base + delta1
			output[fromIdx] = runes[idx1]
			fromIdx++
		}
	} else {
		// 2個目の要素以降。
		for j := 0; j < keyFrames - 1; j++ {
			base := line + delta0 * j
			idx1 := base + delta1
			idx2 := idx1 + delta2

			output[fromIdx] = runes[idx1]
			fromIdx++
			output[fromIdx] = runes[idx2]
			fromIdx++
		}
	}
	if extra > 0 {
		base := line + delta0 * (keyFrames - 1)
		idx1 := base + delta1
		output[fromIdx] = runes[idx1]
	}

	wg.Done()  // 完了したのでwgをデクリメント
}
