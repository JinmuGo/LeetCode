/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    first := head
    arr := make([]int, 0)
    n := 0

    for head != nil {
        arr = append(arr, head.Val)
        n++
        head = head.Next
    }

    head = first

    left, right := 0, n - 1
    maxVal := 0

    for left < right {
        val := arr[left] + arr[right]

        if maxVal < val {
            maxVal = val
        }
        left++
        right--
    }

    return maxVal
}