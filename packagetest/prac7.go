package main

import "fmt"

func do() (int, interface{}){
	return 5, nil
}
func main() {
	result, err := do()
	if (err != nil){
		fmt.Println("えらー")
	}else {
		fmt.Println(result)
	}
}