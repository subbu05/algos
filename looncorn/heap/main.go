package main

type heaps struct {
	array []int
}

func getLeftChild(this *heaps, i int) int {
	lc := 2*i + 1
	if lc > len(this.array) {
		return -1
	}
	return lc
}

func getRightChild(this *heaps, i int) int {
	rc := 2*i + 1
	if rc > len(this.array) {
		return -1
	}
	return rc
}

func getParent(this *heaps, i int) int {
	if i < 0 || i > len(this.array) {
		return -1
	}
	return (i - 1) / 2
}

func siftDown(this *heaps, i int) {
	lc := getLeftChild(this, i)
	rc := getRightChild(this, i)
	index := -1

	if lc != -1 && rc != -1 {
		if this.array[lc] < this.array[rc] {
			index = lc
		} else {
			index = rc
		}
	} else if lc != -1 {
		index = lc
	} else {
		index = rc
	}

	if index == -1 {
		return
	}

	if this.array[index] > this.array[i] {
		this.array[index], this.array[lc] = this.array[lc], this.array[index]
		siftDown(this, index)
	}
}

func siftUp(this *heaps, index int) {
	parentIndex := getParent(this, index)
	if parentIndex != -1 && this.array[index] < this.array[parentIndex] {
		this.array[parentIndex], this.array[index] = this.array[index], this.array[parentIndex]
		siftUp(this, parentIndex)
	}
}

func insert(this *heaps, i int) {
	this.array = append(this.array, i)
	siftUp(this, i)
}

func remove(this *heaps) int {
	res := this.array[0]
	this.array[0] = this.array[len(this.array)-1]
	this.array = this.array[0 : len(this.array)-1]
	siftDown(this, 0)
	return res
}

func peek(this *heaps) int {
	if len(this.array) > 0 {
		return this.array[0]
	}
	return -1
}

func main() {

}
