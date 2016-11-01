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

	n1 := l1
	n2 := l2
	sum := 0
	value := 0
	carry := 0

	for ;; {
		// check exit condition
		if n1 == nil && n2 == nil {
			break
		}

		// increment & fetch current value
		v1 := 0
		if n1 != nil {
			v1 = n1.Val
			n1 = n1.Next
		}
		v2 := 0
		if n2 != nil {
			v2 = n2.Val
			n2 = n2.Next
		}

		// calc sum
		sum = v1 + v2 + carry
		carry = sum/10
		if carry > 0 {
			value = sum - 10
		} else {
			value = sum
		}

		// set current pointer
		current := new(ListNode)
		current.Val = value
		if prev != nil {
			prev.Next = current
		} else {
			ret = current
		}
		prev = current
	}

	if carry > 0 {
		current := new(ListNode)
		current.Val = 1
		prev.Next = current
	}

	return ret;
}
