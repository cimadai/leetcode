import org.junit.Test;
import static org.junit.Assert.*;

public class SolutionTest {
    @Test
    public void testAddTwoNumbers() {
        final ListNode l1 = new ListNode(2);
        l1.next = new ListNode(4);
        l1.next.next = new ListNode(3);

        final ListNode l2 = new ListNode(5);
        l2.next = new ListNode(6);
        l2.next.next = new ListNode(4);

        final Solution sol = new Solution();
        final ListNode ret = sol.addTwoNumbers(l1, l2);
        assertEquals(ret.val, 7);
        assertEquals(ret.next.val, 0);
        assertEquals(ret.next.next.val, 8);
    }
}

