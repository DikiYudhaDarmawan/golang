//package main

import "fmt"

func logging(){
	fmt.Println("Finish runing he program")
}

func runApplication(){
	fmt.Println("Run application")
}

func main(){
	defer logging()
	runApplication()
}