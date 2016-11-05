package q4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	total := len1 + len2

	isOdd := (total & 1) == 1

	var med0 int
	var med1 int
	var l int
	if isOdd {
		med0 = total/2
		med1 = -1
		l = med0
	} else {
		med1 = total/2
		med0 = med1 - 1
		l = med1
	}

	sum := 0
	idx1 := 0
	idx2 := 0
	for i := 0; i < l + 1; i++ {
		n1 := -1
		if idx1 < len1 {
			n1 = nums1[idx1]
		}
		n2 := -1
		if idx2 < len2 {
			n2 = nums2[idx2]
		}

		var n int
		if n1 < 0 {
			idx2++
			n = n2
		} else if n2 < 0 {
			idx1++
			n = n1
		} else if n1 < n2 {
			idx1++
			n = n1
		} else {
			idx2++
			n = n2
		}

		if med0 == i || med1 == i {
			sum += n
		}
	}

	if isOdd {
		return float64(sum)
	} else {
		return float64(sum) / 2
	}
}