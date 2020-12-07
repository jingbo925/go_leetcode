package tree

/* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return false
	}

	if p == nil && q == nil {
		return true
	}

	return q.Val == p.Val && isSameTree(q.Left, p.Left) && isSameTree(q.Right, p.Right)

}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(invertTree(root.Left), root.Right)
}
