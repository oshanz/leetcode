/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	expectSum := targetSum - root.Val

	if expectSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	if hasPathSum(root.Left, expectSum) {
		return true
	}

	if hasPathSum(root.Right, expectSum) {
		return true
	}
	return false
}
