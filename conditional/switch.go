//package main

import "fmt"

func main(){
	n := 5

	switch {
	case n%2 == 0:
		fmt.Println("event number!")
	case n%2 !=0:
		fmt.Println("Odd number!")
	default:
		fmt.Println("Never here!")
	}
}