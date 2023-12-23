/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    was := make(map[ListNode]bool)

    for head != nil {
        if was[*head] {
            return true
        }
        was[*head] = true
        head = head.Next
    }

    return false
}