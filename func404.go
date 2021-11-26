package main

import "fmt"

func rectangle(l int, w int) (area int, parameter int){
	parameter = 2 * (l + w)
	area = l * w
	return
}

func main() {
	var a, p int
	a, p = rectangle(20, 30)
	fmt.Println(a)
	fmt.Println(p)
}