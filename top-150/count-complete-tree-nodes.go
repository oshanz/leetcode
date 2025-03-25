/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    
    if root == nil{
        return 0
    }

    if root.Left == nil && root.Right == nil{
        return 1
    }

    leftNodes := countNodes(root.Left)
    rightNodes := countNodes(root.Right)

    return leftNodes + rightNodes + 1
}
