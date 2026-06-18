/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    if root == nil {
        return 0
    }

    queue := []*TreeNode{root}
    curLevel := 1
    level := 1
    max := root.Val

    for len(queue) > 0 {
        n := len(queue)
        sum := 0
        nextQueue := make([]*TreeNode, 0, n * 2)

        for _, node := range queue {
            sum += node.Val
            if node.Left != nil {
                nextQueue = append(nextQueue, node.Left)
            }
            if node.Right != nil {
                nextQueue = append(nextQueue, node.Right)
            }
        }

        if sum > max {
            max = sum
            level = curLevel
        }

        queue = nextQueue
        curLevel++
    }

    return level
}