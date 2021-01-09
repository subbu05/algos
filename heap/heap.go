package main

type binHeap struct {
	max   int
	array []int
	count int
}

var count int
var heap []int

func main() {

}

func getLeftChild(i int) int {
	left := 2*i + 1
	if left >= count {
		return -1
	}
	return left
}

func getRightChild(i int) int {
	right := 2*i + 2
	if right >= count {
		return -1
	}
	return right
}

func getParent(i int) int {
	if i < 0 || i > count {
		return -1
	}
	return (i - 1) / 2
}

func siftDown(i int) {
	si := -1
	l := getLeftChild(i)
	r := getRightChild(i)

	if l != -1 && r != -1 {
		if heap[l] < heap[r] {
			si = l
		} else {
			si = r
		}
	} else if l == -1 {
		si = r
	} else {
		si = l
	}

	if si == -1 {
		return
	}

	heap[i], heap[si] = heap[si], heap[i]
	siftDown(si)
}

func siftUp(i int) {
	p := getParent(i)
	if p != -1 && heap[p] < heap[i] {
		heap[p], heap[i] = heap[i], heap[p]
		siftUp(i)
	}
}
