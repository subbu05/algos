package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC O(n)
// SC O(n)
func isSymmetricIterative(root *TreeNode) bool {
	q := []*TreeNode{root, root}
	for len(q) > 0 {
		c1 := q[0]
		c2 := q[1]
		q = q[2:] //pop 2 elements
		if c1 == nil && c2 == nil {
			continue
		}
		if c1 == nil || c2 == nil {
			return false
		}
		if c1.Val != c2.Val {
			return false
		}
		q = append(q, c1.Left, c2.Right, c1.Right, c2.Left)
	}
	return true
}

// TC O(n)
// SC O(n)
func isSymmetricRecursive(root *TreeNode) bool {
	return ismirror(root, root)
}

func ismirror(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 != nil || r2 != nil {
		return false
	}
	return r1.Val == r2.Val && ismirror(r1.Left, r2.Right) && ismirror(r1.Right, r2.Left)
}

func main() {

}
