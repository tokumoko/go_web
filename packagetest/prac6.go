package main

import "fmt"

func div(a, b int) (int, int) {
	p := a / b
	q := a * b
	return p, q
}

func main() {
	var a interface{}
	fmt.Println(a)
	fmt.Println(div(10, 2))
}