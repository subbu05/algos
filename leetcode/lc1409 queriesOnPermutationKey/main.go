package main

func main() {

}

func processQueries(queries []int, m int) []int {
	P := []int{}
	res := []int{}

	for i := 1; i <= m; i++ {
		P = append(P, i)
	}

	for i := 0; i < len(queries); i++ {
		for ind, v := range P {
			if queries[i] == v {
				res = append(res, ind)

				copy(P[1:ind+1], P[:ind])
				P[0] = v
			}
		}
	}

	return res
}
