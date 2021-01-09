package main

import "fmt"

type graphMatrix struct {
	nodes [][]int
	isDirected bool
}

func main() {

	graph := &graphMatrix{isDirected: true,}
	for i:=0;i<5;i++{
		graph.nodes = append(graph.nodes,make([]int,5))
	}


	addEdge(1,2,graph)
	addEdge(3,4, graph)
	addEdge(4,2, graph)
	addEdge(3,2, graph)

	fmt.Println(topSort(graph))

}


func addEdge(v1,v2 int, graph *graphMatrix) {
	graph.nodes[v1][v2] = 1
	if graph.isDirected == false {
		graph.nodes[v2][v1] = 1
	}
}

func getInDegree(graph *graphMatrix, v int) int {
	res := 0
	for i:=0;i<len(graph.nodes);i++{
		if graph.nodes[i][v] == 1 {
			res++
		}
	}
	return res
}

func getAdjacentNodes(graph *graphMatrix, v int) []int {
	res := []int{}
	for i:=0;i<len(graph.nodes);i++{
		if graph.nodes[v][i]==1{
			res = append(res, i)
		}
	}
	return res
}

func topSort(graph *graphMatrix) []int {
	m := make(map[int]int)
	q := []int{}
	for i:=0;i<len(graph.nodes);i++{
		m[i]= getInDegree(graph,i)
		if m[i] == 0 {
			q = append(q, i)
		}
	}
	res := []int{}
	for len(q) > 0 {
		//fmt.Println(m)
		cur := q[0]
		q = q[1:]
		res = append(res,cur)
		adjNodes := getAdjacentNodes(graph,cur)
		delete(m,cur)
		for _, node := range adjNodes {
			m[node] -= 1
			if m[node]==0 {
				q=append(q,node)
			}
		}
	}

	if len(m) > 0 {
		return []int{}
	}

	return res

}