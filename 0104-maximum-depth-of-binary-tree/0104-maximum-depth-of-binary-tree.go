/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    return dfs(root)
}

func dfs(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := dfs(root.Left)
    right := dfs(root.Right)

    return max(left, right) + 1
}