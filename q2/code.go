package q2

type ListNode struct {
    Val int
    Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret *ListNode
	var prev *ListNode

	next := 0

	for ;; {
		var current *ListNode = nil
		sum := 0

		// increment & fetch current value
		if l1 != nil {
			// reuse
			current = l1
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			// reuse
			current = l2
			sum += l2.Val
			l2 = l2.Next
		}

		// check exit condition
		if current == nil {
			break
		}

		// calc sum
		sum += next
		if sum > 9 {
			next = 1
			current.Val = sum - 10
		} else {
			next = 0
			current.Val = sum
		}

		// set current pointer
		if prev != nil {
			prev.Next = current
		} else {
			ret = current
		}
		prev = current
	}

	if next == 1 {
		prev.Next = new(ListNode)
		prev.Next.Val = 1
	}

	return ret;
}
