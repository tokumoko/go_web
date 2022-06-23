package main

import(
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3}
	fmt.Println(a)
	b := [5]int{}
	fmt.Println(b)
	ia := [3]int{}
	fmt.Println(ia)
	ua := [3]uint{}
	fmt.Println(ua)
	ba := [3]bool{}
	fmt.Println(ba)
	fa := [3]float64{}
	fmt.Println(fa)
	ca := [3]complex128{}
	fmt.Println(ca)
	ra := [3]rune{}
	fmt.Println(ra)
	sa := [3]string{}
	fmt.Println(sa)
}