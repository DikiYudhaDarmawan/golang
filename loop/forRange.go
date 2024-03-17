//package main

import "fmt"

func main(){
	names := []string{"Jonh","Paul","George","Ringo"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}