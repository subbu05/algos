package main

type node struct {
	Val  int
	Next *node
}

func main() {
	var head *node = nil

}

func detectCycle(head *node) bool {
	m := make(map[*node]bool)
	cur := head
	for cur != nil {
		if _, ok := m[cur]; ok {
			return true
		}
		m[cur] = true
		cur = cur.Next
	}
	return false
}

func detectCycleFloyd(head *node) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}
