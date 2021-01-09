/**
 * Definition for a binary tree node.*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res = []int{}

func increasingBST(root *TreeNode) *TreeNode {
	res = []int{}
	ans := &TreeNode{}

	cur := ans
	helper(root)
	for _, v := range res {
		cur.Right = &TreeNode{
			Val: v,
		}
		cur = cur.Right
	}
	return ans.Right
}

func helper(root *TreeNode) {
	if root != nil {
		helper(root.Left)
		res = append(res, root.Val)
		helper(root.Right)
	}
}
