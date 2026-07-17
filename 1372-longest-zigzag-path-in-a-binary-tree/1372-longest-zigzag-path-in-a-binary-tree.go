/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 // direction true -> left, false -> right
func longestZigZag(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := maxZigZag(root.Left, false)
    right := maxZigZag(root.Right, true)
    val1 := longestZigZag(root.Left)
    val2 := longestZigZag(root.Right)

    return max(left, right, val1, val2)
}

func maxZigZag(node *TreeNode, direction bool) int {
    if node == nil {
        return 0
    }
    cnt := 0

    for node != nil {
        if direction {
            // left
            node = node.Left
        } else {
            node = node.Right
        }
        cnt++
        direction = !direction
    }

    return cnt
}