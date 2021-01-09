package main

const MaxNodes = 10

type graph struct {
	dag    bool
	matrix [MaxNodes][MaxNodes]int
}

var g graph

func main() {
	g.matrix = [MaxNodes][MaxNodes]int{}
	g.dag = false
}

func addEdge(v1, v2 int) {
	if isValidVertex(v1) && isValidVertex(v2) {
		g.matrix[v1][v2] = 1
		if g.dag == false {
			g.matrix[v2][v1] = 1
		}
	}
}

func getAdjVertices(v int) []int {
	res := []int{}
	if isValidVertex(v) {
		for i := 0; i < MaxNodes; i++ {
			if g.matrix[v][i] == 1 {
				res = append(res, i)
			}
		}
	}
	return res
}

func isValidVertex(v int) bool {
	return v >= 0 && v < MaxNodes
}
