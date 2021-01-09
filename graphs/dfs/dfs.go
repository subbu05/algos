package main

import "fmt"

type graph struct {
	matix   [5][5]int
	visited [5]int
}

var g graph

func main() {
	g := &graph{
		matix:   [5][5]int{},
		visited: [5]int{},
	}
	addEdge(g, 0, 4)
	//g.matix[0] = append(g.matix[0],0,1,1,1,0)
	fmt.Println(*g)
}

func addEdge(g *graph, v1, v2 int) {
	if isValid(v1) && isValid(v2) {
		g.matix[v1][v2] = 1
		g.matix[v2][v1] = 1
	}
}

func dfs(v int) {
	if isValid(v) && g.visited[v] == 0 {
		fmt.Println(v)
		g.visited[v] = 1
		for i := 0; i < len(g.matix); i++ {
			if g.matix[v][i] == 1 && g.visited[i] == 0 {
				g.visited[i] = 1
				fmt.Println(i)
			}
		}
	}
}

func bfs(v int) {
	if isValid(v) {
		q := []int{v}

		for len(q) > 0 {
			cur := q[0]
			fmt.Println(cur)
			g.visited[cur] = 1
			q = q[1:]
			for i := 0; i < len(g.matix); i++ {
				if g.matix[v][i] == 1 {
					q = append(q, i)
				}
			}
		}
	}
}

func isValid(v int) bool {
	return v >= 0 && v < len(g.matix)
}
