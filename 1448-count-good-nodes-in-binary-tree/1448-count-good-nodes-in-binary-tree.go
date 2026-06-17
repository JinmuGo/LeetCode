/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    cnt := 1

    isGood(root.Val, &cnt, root.Left) 
    isGood(root.Val, &cnt, root.Right) 

    return cnt
}

func isGood(maxVal int, cnt *int, root *TreeNode) {
    if root == nil {
        return 
    }

    if maxVal <= root.Val {
        maxVal = root.Val
        (*cnt)++
    }

    isGood(maxVal, cnt, root.Left)
    isGood(maxVal, cnt, root.Right)
}