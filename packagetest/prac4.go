package main

import "fmt"

func main() {
	r := '\xff'
	fmt.Printf("%c\n", r)
	t := '㌦'
	fmt.Printf("%x\n", t)
	s := "golang"
	fmt.Printf("%v\n", s)
}