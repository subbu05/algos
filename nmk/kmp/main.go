package main

func main() {

}

func buildPrefixTable(s string) []int {
	F := make([]int, len(s))
	i := 1
	j := 0
	F[0] = 0

	for i < len(s) {
		if s[i] == s[j] {
			F[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = F[j-1]
		} else {
			F[i] = 0
			i++
		}
	}
	return F
}
