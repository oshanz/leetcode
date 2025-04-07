/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	return slices.Equal(leafs(root1), leafs(root2))
}

func leafs(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}
	lv := leafs(node.Left)
	rv := leafs(node.Right)
	return append(lv, rv...)
}
