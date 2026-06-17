/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    arr := []int{}

    dfs(&arr, root1)

    return compare(&arr, root2) && len(arr) == 0
}

func dfs(arr *[]int, root *TreeNode) {
    if root == nil {
        return
    }

    if root.Left == nil && root.Right == nil {
        *arr = append(*arr, root.Val)
        return
    }

    dfs(arr, root.Left)
    dfs(arr, root.Right)
}

func compare(arr *[]int, root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    if root.Left == nil && root.Right == nil {
        if (len(*arr) == 0) {
            return false
        }

        elem := (*arr)[0]
        if elem != root.Val {
            return false
        } 
        
        *arr = (*arr)[1:]
        return true
    }

    return compare(arr, root.Left) && compare(arr, root.Right)
}