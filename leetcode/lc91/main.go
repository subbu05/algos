package main

func main() {
	//m := buildMap()

}

func buildMap() map[byte]string {
	m := make(map[byte]string)
	for i := 1; i <= 26; i++ {
		m[byte(i)] = string(i + 64)
	}
	return m
}
