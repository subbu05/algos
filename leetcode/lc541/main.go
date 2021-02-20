package main

func main() {
	reverseStr("mebeddqsgl", 3)
}

func reverseStr(s string, k int) string {
	r := []rune(s)
	for i := 0; i < len(r); {
		if i+k-1 < len(r) {
			reverse(r, i, i+k-1) // next iteration after 2K characters.
		} else {
			reverse(r, i, len(r)-1) // leftover reverse.
		}
		i = i + 2*k // skip 2K chars for every run.
	}
	return string(r)
}

func reverse(r []rune, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}
