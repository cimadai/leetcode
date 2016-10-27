package q1

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, e := range nums {
		if x, found := m[e]; found {
			return []int{x, i}
		}
		m[target - e] = i
	}
	return []int{}
}
