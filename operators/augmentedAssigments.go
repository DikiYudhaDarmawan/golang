//package main

import (
	"fmt"
)

func main(){
	a,b := 2,3

	a += b
	fmt.Println(a)

	b -= 2
	fmt.Println(b)

	b *= 10
	fmt.Println(b)

	b /= 5
	fmt.Println(b)

	a %= 3
	fmt.Println(a)
}