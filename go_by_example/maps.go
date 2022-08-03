package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	delete(m, "k2")

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

}
