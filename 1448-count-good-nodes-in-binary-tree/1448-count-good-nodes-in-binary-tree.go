/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    return dfs(root, root.Val)
}

func dfs(root *TreeNode, maxVal int) int {
    if root == nil {
        return 0
    }

    count := 0

    if root.Val >= maxVal {
        count++
        maxVal = root.Val
    }

    return count + dfs(root.Left, maxVal) + dfs(root.Right, maxVal)
}

