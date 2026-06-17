/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    cnt := 0

    var isGood func(maxVal int, root *TreeNode)

    isGood = func(maxVal int, root *TreeNode) {
        if root == nil {
            return 
        }
    
        if maxVal <= root.Val {
            maxVal = root.Val
            cnt++
        }
    
        isGood(maxVal, root.Left)
        isGood(maxVal, root.Right)
    }

    isGood(root.Val, root)

    return cnt
}

