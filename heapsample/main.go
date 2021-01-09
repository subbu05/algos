package main

type heap struct {
	elements []int
	MAX      int
	count    int
}

func new(max int) heap {
	return heap{
		elements: make([]int, max),
		MAX:      max,
		count:    -1,
	}
}

func getLeftChild(h heap, v int) int {
	if v < 0 || v > h.MAX {
		return -1
	}
	return 2*v + 1
}

func getRightChild(h heap, v int) int {
	if v < 0 || v > h.MAX {
		return -1
	}
	return 2*v + 2
}

func getParent(h heap, v int) int {
	if v < 0 || v > h.MAX {
		return -1
	}
	return (v - 1) / 2
}
func siftDown(h heap, i int) {
	leftIndex := getLeftChild(h, i)
	rightIndex := getRightChild(h, i)
	min := -1

	if leftIndex != -1 && rightIndex != -1 {
		if leftIndex < rightIndex {
			min = getMin(leftIndex, min)
		} else {
			min = getMin(rightIndex, min)
		}
	} else if leftIndex != -1 {
		min = getMin(leftIndex, min)
	} else if rightIndex != -1 {
		min = getMin(rightIndex, min)
	}
	if min == -1 {
		return
	}

	if min < i {
		h.elements[min], h.elements[i] = h.elements[i], h.elements[min]
		siftDown(h, min)
	}
}

func siftUp(h heap, i int) {
	pi := getParent(h, i)
	if pi != -1 && h.elements[pi] > h.elements[i] {
		h.elements[pi], h.elements[i] = h.elements[i], h.elements[pi]
		siftUp(h, pi)
	}
}

func insertIntoHeap(h heap, v int) {
	if h.count < len(h.elements) {
		h.elements[h.count] = v
		siftUp(h, h.count)
		h.count++
	}
}

func removeFromHeap(h heap) {
	if h.count >= 0 {
		h.elements[0] = h.elements[h.count-1]
		h.count--
		siftDown(h, 0)
	}
}

func getMaxFromMinHeap(h heap) int {
	lastIndex := h.count - 1
	parent := getParent(h, lastIndex)
	firstLeaf := parent + 1
	max := h.elements[firstLeaf]
	for i := firstLeaf; i <= lastIndex; i++ {
		if h.elements[i] > max {
			max = h.elements[i]
		}
	}
	return max
}

func getMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
