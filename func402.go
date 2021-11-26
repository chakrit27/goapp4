package main

import "fmt"

func fun() int{
	a := 10
	b := 20
	msg := "Welcome"
	fmt.Println(msg)
	add := a + b
	return add
}

func main() {
	x := fun()
	fmt.Print(x)
}