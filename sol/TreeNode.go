package sol

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	return IsEqual(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func IsEqual(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	if subRoot != nil {
		if root.Val != subRoot.Val {
			return false
		}
		return IsEqual(root.Left, subRoot.Left) && IsEqual(root.Right, subRoot.Right)
	} else {
		return false
	}
}
