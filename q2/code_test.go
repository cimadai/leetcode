package q2

import (
	"testing"
)

func createTestData() (*ListNode, *ListNode) {
	l1 := new(ListNode)
	l1.Val = 2
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 3

	l2 := new(ListNode)
	l2.Val = 5
	l2.Next = new(ListNode)
	l2.Next.Val = 6
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4

	return l1, l2
}

func TestCode(t *testing.T) {
	l1, l2 := createTestData()

	ret := addTwoNumbers(l1, l2)

	if ret.Val != 7 {
		t.Error("First value must be 7.")
	}

	if ret.Next.Val != 0 {
		t.Error("Second value must be 0.")
	}

	if ret.Next.Next.Val != 8 {
		t.Error("Last value must be 8.")
	}

	if ret.Next.Next.Next != nil {
		t.FailNow()
		t.Error("This test must produce 3 nodes.")
	}
}

func BenchmarkCode(b *testing.B) {
	l1, l2 := createTestData()

	for i := 0; i < b.N; i++ {
		addTwoNumbers(l1, l2)
	}
}