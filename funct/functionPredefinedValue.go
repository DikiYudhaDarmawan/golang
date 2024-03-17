//package main

import "fmt"

func getName() (firstName, lastName string) {
	return "diki", "darmawan"
}

func main(){
	x,y := getName()
	fmt.Println(x, y)
}