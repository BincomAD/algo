/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    answ := head

    for head != nil {
        for head.Next != nil && head.Next.Val == head.Val {
            head.Next = head.Next.Next
        }
        head = head.Next
    }

    return answ
}