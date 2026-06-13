/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head.Next == nil {
        return nil
    }
    n := 0
    first := head

    for head != nil {
        n++
        head = head.Next
    }
    head = first

    for i := 0; i < n / 2 - 1; i++ {
        head = head.Next
    }

    if head.Next != nil && head.Next.Next != nil {
        head.Next = head.Next.Next
    } else {
        head.Next = nil
    }
    return first
}