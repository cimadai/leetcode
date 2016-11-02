/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
public class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode ret = null;
        ListNode prev = null;
        ListNode current;

        ListNode n1 = l1;
        ListNode n2 = l2;

        Integer v1, v2, sum, value, carry = 0;
        do {
            // increment & fetch current value
            if (n1 == null) {
                v1 = 0;
            } else {
                v1 = n1.val;
                n1 = n1.next;
            }

            if (n2 == null) {
                v2 = 0;
            } else {
                v2 = n2.val;
                n2 = n2.next;
            }

            // calc sum
            sum = v1 + v2 + carry;
            carry = sum/10;
            value = carry > 0 ? sum - 10 : sum;

            // set current pointer
            current = new ListNode(value);
            if (prev != null) {
                prev.next = current;
            } else {
                ret = current;
            }
            prev = current;

        } while (n1 != null || n2 != null);

        if (carry > 0) {
            prev.next = new ListNode(1);
        }
        return ret;
    }
}
