package main

import "fmt"

type graph struct {
	elements [][]int
}

/*func main() {
    temp := [][]int{}
    for i:=0;i<5;i++{
        temp = append(temp, make([]int,5))
    }
    g := graph {
        elements : temp,
    }

    addEdge(g,0,1)
    addEdge(g,0,2)
    addEdge(g,2,3)
    addEdge(g,3,4)
    addEdge(g,1,4)

    view(g)

    m := make(map[int]int)
    q := []int{}
    for i:=0;i<len(g.elements);i++{
        m[i] = getInDegreeVertex(g,i)
        if m[i] == 0 {
            q = append(q,i)
        }
    }


}


func addEdge(g graph, v1,v2 int) {
    g.elements[v1][v2]=1
}

func view(g graph) {
    for i:=0;i<len(g.elements);i++{
        fmt.Println()
        for j:=0;j<len(g.elements[0]);j++{
            fmt.Print(g.elements[i][j])
        }
    }
    println()
}

func getInDegree(g graph) []int {
    res := make([]int,len(g.elements))
    sum := 0
    for i:=0;i<len(g.elements);i++{
        sum  = 0
        for j:=0;j<len(g.elements);j++{
            if g.elements[i][j]==1{
                sum++
            }
        }
        res[i] = sum
    }
    return res
}

func getInDegreeVertex(g graph, n int) int {
    res := 0
    sum := 0
    for i:=0;i<len(g.elements);i++{
            if g.elements[n][i]==1{
                sum++
            }

    }
    return res
}
*/

func addEdge(g graph, v1, v2 int) {
	if isValid(v1) && isValid(v2) {

	}
}

func isValid(v int) bool {
	return true
}
