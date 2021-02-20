package main

import "fmt"

type node struct {
	key   int
	value int
	next  *node
}

type hash struct {
	elements []*node
	size     int
}

func put(h *hash, k, v int) {
	if h != nil {
		temp := h.elements[k%h.size]
		t := &node{
			key:   k,
			value: v,
			next:  temp,
		}
		h.elements[k%h.size] = t
	}
}

func constructor(n int) *hash {
	return &hash{
		elements: make([]*node, n),
		size:     n,
	}
}

func main() {
	h := constructor(7)

	put(h, 1, 5)
	v, ok := get(h, 1)
	if ok == false {
		fmt.Println("Not found")
	} else {
		fmt.Println(v)
	}

}

func get(h *hash, k int) (int, bool) {
	temp := h.elements[k%h.size]
	return isExists(temp, k)
}

func isExists(head *node, k int) (int, bool) {
	if head == nil {
		return 0, false
	}
	cur := head
	for cur != nil {
		if cur.key == k {
			return cur.value, true
		}
		cur = cur.next
	}
	return 0, false
}
