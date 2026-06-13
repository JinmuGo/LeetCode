/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    behindTheFirst := &ListNode{0, head}

    bridge, oneStep, twoStep := behindTheFirst, head, head

    for twoStep != nil && twoStep.Next != nil {
        bridge = bridge.Next
        oneStep = oneStep.Next
        twoStep = twoStep.Next.Next
    }

    bridge.Next = oneStep.Next

    return behindTheFirst.Next
}   