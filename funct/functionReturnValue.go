//package main

import "fmt"

func gethelloworld() (string, string){
	return "hello", "world"
}

func main(){
	x,y := gethelloworld()
	fmt.Println(x, y)
}