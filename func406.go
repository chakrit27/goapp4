package main

import "fmt"

func fectorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * fectorial(num-1)
}

func main() {
	fmt.Println(fectorial(5))
}