/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {
    var carry = 0, ret = [];

    for (;;) {
        var isExit = true;
        var sum = 0;
        if (l1) {
            isExit = false;
            sum += l1.val|0;
            l1 = l1.next;
        }
        if (l2) {
            isExit = false;
            sum += l2.val|0;
            l2 = l2.next;
        }
        if (isExit) break;

        sum += carry;
        if (sum > 9) {
            carry = 1;
            ret[ret.length] = sum - 10;
        } else {
            carry = 0;
            ret[ret.length] = sum;
        }
    }

    if (carry) { ret[ret.length] = 1; }

    return ret;
};
