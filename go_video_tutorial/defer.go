package main

import "fmt"

func add(x, y int) (z1, z2 int) {
	defer fmt.Println("hello")
	z1 = x + y
	z2 = x - y
	return
}

func anotherExample() string {
	defer fmt.Println("hello")
	fmt.Println("world")
	return "HELLO"
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(anotherExample())

}
