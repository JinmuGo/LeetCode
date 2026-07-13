/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	var state int

	if root.Left != nil {
		state |= 2
	}

	if root.Right != nil {
		state |= 1
	}

    switch state {
        case 0:
            return nil
        case 1:
            return root.Right
        case 2:
            return root.Left

        default:
            minNode := root.Right
            for minNode.Left != nil {
                minNode = minNode.Left
            }
            root.Val = minNode.Val
            root.Right = deleteNode(root.Right, minNode.Val)
            return root
    }
}