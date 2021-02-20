package main

import "fmt"

type graph struct {
	elements [][]int
}

func main() {
	temp := [][]int{}
	for i := 0; i < 5; i++ {
		temp = append(temp, make([]int, 5))
	}
	g := &graph{
		elements: temp,
	}

	addEdge(g, 0, 1)
	addEdge(g, 0, 2)
	addEdge(g, 0, 4)
	addEdge(g, 2, 4)
	addEdge(g, 4, 3)
	//fmt.Printf("%#v\n",g)
	visited := make([]int, 5)
	/* for i:=0;i<5;i++{
	    dfs(g,i,visited)
	}*/
	bfs(g, 0, visited)
}

func addEdge(g *graph, v1, v2 int) {
	if v1 < 0 || v1 >= len(g.elements) || v2 < 0 || v2 >= len(g.elements) {
		return
	}
	g.elements[v1][v2] = 1
	g.elements[v2][v1] = 1
}

func dfs(g *graph, n int, visited []int) {
	if visited[n] == 1 {
		return
	}
	visited[n] = 1
	r := getAdjacent(g, n)
	//fmt.Printf("%#v\n",r)
	for _, v := range r {
		dfs(g, v, visited)
	}
	fmt.Println(n)
}

func getAdjacent(g *graph, n int) []int {
	res := []int{}
	for i := 0; i < len(g.elements); i++ {
		if g.elements[n][i] == 1 {
			res = append(res, i)
		}
	}
	return res
}

func bfs(g *graph, n int, visited []int) {

	q := []int{n}
	for len(q) > 0 {
		n = q[0]
		q = q[1:]
		if visited[n] == 1 {
			continue
		}

		fmt.Println(n)
		visited[n] = 1

		r := getAdjacent(g, n)
		for _, v := range r {
			q = append(q, v)
		}
	}
}
