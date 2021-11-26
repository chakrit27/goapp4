package main

import "fmt"

func area(length, width int) int {
	Ar := length * width
	return Ar
}
func main() {
	a := area(12, 10)
	fmt.Printf("Area of rectsngle is: %v\n",a)
	fmt.Println("Area of rectsngle is:",a)
}