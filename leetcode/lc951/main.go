package main

type node struct {
	Val   int
	Left  *node
	Right *node
}

func main() {

}

// TC O(min(L1,L2)) SC(min(H1,H2))
func flipEquivRecursive(r1, r2 *node) bool {
	if r1 == r2 {
		return true
	}

	if r1 == nil || r2 == nil || r1.Val != r2.Val {
		return false
	}
	return flipEquivCanonical(r1.Left, r2.Left) && flipEquivCanonical(r1.Right, r2.Right) || flipEquivCanonical(r1.Left, r2.Right) && flipEquivCanonical(r1.Right, r2.Left)
}

// TC O(L1+L2)) SC O(L1+L2)
func flipEquivCanonical(root1 *node, root2 *node) bool {

	if root1 == root2 {
		return true
	}

	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}

	v1 := []int{}
	v2 := []int{}

	dfs(root1, &v1)
	dfs(root2, &v2)

	if len(v1) != len(v2) {
		return false
	}

	return isEqual(v1, v2)

}

func dfs(root *node, v *[]int) {
	if root != nil {
		*v = append(*v, root.Val)
		L := -1
		R := -1
		if root.Left != nil {
			L = root.Left.Val
		}
		if root.Right != nil {
			R = root.Right.Val
		}
		if L < R {
			dfs(root.Left, v)
			dfs(root.Right, v)
		} else {
			dfs(root.Right, v)
			dfs(root.Left, v)
		}
		*v = append(*v, 0)
	}
}

func isEqual(v1, v2 []int) bool {
	for i := 0; i < len(v1); i++ {
		if v2[i] != v1[i] {
			return false
		}
	}
	return true
}
