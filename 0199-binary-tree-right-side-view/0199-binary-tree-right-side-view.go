/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil {
        return nil
    }
   queue := []*TreeNode{root}
   arr := make([]int, 0)

    for len(queue) > 0 {
        L := len(queue)

        for i := 0; i < L; i++ {
            node := queue[0]
            queue = queue[1:]

            if i == L - 1 {
                arr = append(arr, node.Val)
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }

    return arr
}