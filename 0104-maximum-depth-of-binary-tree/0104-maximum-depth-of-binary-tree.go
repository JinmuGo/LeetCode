/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    return dfs(root, 0)
}

func dfs(root *TreeNode, depth int) int {
    if root == nil {
        return depth
    }

    left := dfs(root.Left, depth + 1)
    right := dfs(root.Right, depth + 1)

    return max(left, right)
}