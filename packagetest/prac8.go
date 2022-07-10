package main

import "fmt"

func doSome() (a int) {
	return
}
func Some1() (x, y int) {
	y = 5
	return
}

func main(){
	fmt.Println(doSome())
	fmt.Println(Some1())
}