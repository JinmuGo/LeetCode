/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    queue := []*TreeNode{root}
    curLevel := 1
    level := 1
    max := root.Val

    for len(queue) > 0 {
        n := len(queue)
        sum := 0

        for i := 0; i < n; i++ {
            node := queue[0]
            queue = queue[1:]

            sum += node.Val

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        if sum > max {
            max = sum
            level = curLevel
        }

        curLevel++
    }

    return level
}