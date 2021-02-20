package main

import "sort"

type graph struct {
	node int
	next *graph
}

type distanceInfo struct {
	distance   int
	lastvertex int
}

func getAdjacentNodes(g *graph) []int {
	res := []int{}
	cur := g
	for cur != nil {
		res = append(res, cur.node)
		cur = cur.next
	}
	sort.Ints(res)
	return res
}

func buildAdjacentMatrix(g *graph, v int) {
	temp := graph{
		node: v,
	}
	if g == nil {
		g = &temp
	} else {
		temp.next = g
		g = &temp
	}
}

// DistanceInfo is function.
func getDistanceInfo() distanceInfo {
	return distanceInfo{
		distance:   -1,
		lastvertex: -1,
	}
}

func spath(gnodes []graph, source int) {
	m := make(map[int]distanceInfo, 0)
	for i := 0; i < 5; i++ {
		m[i] = getDistanceInfo()
	}

	d := m[source]
	d.distance = 0
	d.lastvertex = source
	m[source] = d

	q := []int{source}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		n := getAdjacentNodes(&gnodes[cur])
		for _, v := range n {
			cd := m[v].distance
			if cd == -1 {
				cd = m[cur].distance + 1

			}
		}
	}

}

func main() {
	g := make([]graph, 5)

	for i := 0; i < 5; i++ {
		g[0] = graph{}
	}

	spath(g, 0)
}
