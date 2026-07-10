/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prevHead *ListNode
    var nextHead *ListNode
    
    for head != nil {
        nextHead = head.Next
        head.Next = prevHead

        prevHead = head
        head = nextHead
    }

    return prevHead
}