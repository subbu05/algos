package main

import "fmt"

type tree struct {
	val   byte
	left  *tree
	right *tree
}

func main() {
	fmt.Print(preToIn("DBEAFC", "ABDECF"))
}

func preToIn(pre, in string) *tree {
	if len(pre) == 0 || len(in) == 0 || len(pre) != len(in) {
		return nil
	}
	return helper(pre, 0, len(pre)-1, in, 0, len(in)-1)
}

func helper(pre string, ps, pe int, in string, is, ie int) *tree {
	if ps > pe || is > ie {
		return nil
	}
	offset := is
	fmt.Println(ps+1, pe-ps+offset, is, offset-1)
	cur := &tree{
		val:   pre[ps],
		left:  nil,
		right: nil,
	}

	for ; offset < ie; offset++ {
		if in[offset] == cur.val {
			break
		}
	}

	cur.left = helper(pre, ps+1, ps-is+offset, in, is, offset-1)
	cur.right = helper(pre, ps+offset-is+1, pe, in, offset+1, ie)

	return cur
}
