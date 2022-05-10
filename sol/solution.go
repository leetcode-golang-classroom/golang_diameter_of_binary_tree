package sol

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	FindDiameter(root, &max)
	return max
}
func FindDiameter(root *TreeNode, max *int) int {
	if root == nil {
		return -1 // -1 height
	}
	leftLevel := FindDiameter(root.Left, max)
	rightLevel := FindDiameter(root.Right, max)
	curMax := leftLevel + rightLevel + 2
	*max = Max(*max, curMax)
	return Max(leftLevel, rightLevel) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
