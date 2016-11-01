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
    var n1 = l1,
        n2 = l2,
        carry = 0,
        sum = 0,
        ret = [];

    do {
        sum = ((n1 && n1.val) || 0) + ((n2 && n2.val) || 0) + carry;
        carry = (sum/10)|0;
        ret[ret.length] = carry ? sum - 10 : sum;
        n1 = n1 && n1.next;
        n2 = n2 && n2.next;
    } while (n1 || n2);

    if (carry) { ret.push(1); }

    return ret;
};

