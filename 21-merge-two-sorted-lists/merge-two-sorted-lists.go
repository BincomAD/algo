/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head := &ListNode{}
    next := head

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            next.Next = list1
            list1 = list1.Next
        } else {
            next.Next = list2
            list2 = list2.Next
        }

        next = next.Next
    }

    if list1 == nil {
        next.Next = list2
    } else {
        next.Next = list1
    }

    return head.Next
}