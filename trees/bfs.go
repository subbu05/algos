package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func main() {

}

func bfs(root *Node) {
	if root == nil {
		return
	}

	q := []*Node{root}

	for len(q) > 0 {
		temp := q[0]
		q = q[1:]
		fmt.Println(temp.Data)
		if temp.Left != nil {
			q = append(q, temp.Left)
		}
		if temp.Right != nil {
			q = append(q, temp.Right)
		}
	}
}

func minBst(root *Node) {
	if root == nil {
		return
	}

	if root.Left == nil {
		fmt.Println(root.Data)
		return
	}

	minBst(root.Left)

}

func maxDepth(root *Node) int {

	if root == nil || root.Left == nil && root.Right == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}

func isMirror(root *Node) bool {
	return mHelper(root, root)
}

func mHelper(r1, r2 *Node) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 != nil && r2 == nil || r1 == nil && r2 != nil {
		return false
	}

	return mHelper(r1.Left, r2.Right) && mHelper(r1.Right, r2.Left)

}
