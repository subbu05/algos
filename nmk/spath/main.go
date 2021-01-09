package main

import (
	"fmt"
)

type graph struct {
	nodes [][]int
}

type distanceInfo struct {
	lv int
	distance int
}

func buildTable(g *graph, source int) map[int]*distanceInfo {
	m := make(map[int]*distanceInfo)
	for i:=0;i<len(g.nodes);i++{
		m[i]=&distanceInfo{
			distance: -1,
			lv: -1,
		}
	}

	m[source].distance=0
	m[source].lv=source

	q := []int{source}

	for len(q) > 0 {
		cv := q[0]
		q = q[1:]
		adjNodes := getAdjacentNodes(g,cv)
		for _, v := range adjNodes {
			if m[v].distance == -1 {
				m[v].distance = m[cv].distance + 1
				m[v].lv = cv
				if len(getAdjacentNodes(g,v)) > 0 {
					q = append(q,v)
				}
			}
		}
	}

	return m
}

func getAdjacentNodes(g *graph, v int) []int {
	res := []int {}
	for i:=0;i<len(g.nodes);i++{
			if g.nodes[v][i] == 1 {
				res = append(res, i)
			}
	}
	return res
}

func view(g *graph) {
	for i:=0;i<len(g.nodes);i++ {
		fmt.Println()
		for k:=0;k<len(g.nodes);k++{
			fmt.Printf("%d ",g.nodes[i][k])
		}
	}
}

func main() {
	temp := [][]int{}
	for i:=0;i<5;i++{
		temp =append(temp, make([]int,5))
	}
	g := &graph{nodes: temp }
	addEdge(g,0,1)
	addEdge(g,0,3)
	addEdge(g,0,2)
	addEdge(g,2,3)
	addEdge(g,3,1)
	addEdge(g,1,4)
	addEdge(g,3,4)

	view(g)
	//m:=buildTable(g,3)
	spath(g,3,0)
}


func addEdge(g *graph,s,d int) {
	g.nodes[s][d]=1
}

func spath(g *graph, s, d int) {
	 m := buildTable(g,s)
	 fmt.Println("SPATH is ")
     st := []int{d}

	 pv := m[d].lv

	 for pv != -1 && pv != s {
	 	fmt.Println(pv)
	 	st = append(st, pv)
	 	pv = m[pv].lv
	 }

	 if pv == -1 {
	 	fmt.Println("No Path")
	 } else {
	 	st = append(st,s)
	 	fmt.Println(st)
	 }

     //fmt.Println(st)
}