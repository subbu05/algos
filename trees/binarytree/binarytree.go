package main

import "fmt"

type node struct {
	Val   int
	Left  *node
	Right *node
}

func main() {
	root := &node{
		Val:   0,
		Left:  nil,
		Right: nil,
	}

	root.Left = &node{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	root.Right = &node{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	preOrderRecursive(root)
	fmt.Println(isExists(root, 2))
	inOrderIterative(root)
}

func preOrderIterative(root *node) {
	fmt.Println(root)
}

func inOrderIterative(root *node) {
	if root == nil {
		return
	}
	s := []*node{}
	cur := root
	for len(s) > 0 || cur != nil {
		for cur != nil {
			s = append(s, cur)
			cur = cur.Left
		}
		cur = s[len(s)-1]
		fmt.Println(cur.Val)
		s = s[:len(s)-1]
		cur = cur.Right
	}
}

func preOrderRecursive(root *node) {
	if root != nil {
		fmt.Println(root.Val)
		preOrderRecursive(root.Left)
		preOrderRecursive(root.Right)
	}
}

func getMin(root *node) int {
	if root == nil {
		return -1 << 15
	}
	l := getMin(root.Left)
	r := getMin(root.Right)
	min := r
	if l < r {
		min = l
	}

	if min < root.Val {
		return min
	}
	return root.Val
}

func isExists(root *node, key int) bool {
	if root == nil {
		return false
	}

	if root.Val == key {
		return true
	}

	return isExists(root.Left, key) || isExists(root.Right, key)

}

func isExistsWithChannels(root *node, key int) bool {
	return false
}

func sizeBinaryTree(root *node) int {
	if root == nil {
		return 0
	}

	l := 0
	r := 0

	if root.Left != nil {
		l = sizeBinaryTree(root.Left)
	}

	if root.Right != nil {
		r = sizeBinaryTree(root.Right)
	}

	return 1 + l + r
}
