//package main

import (
	"fmt" 
)

func main(){
	var employees map[string]string = map[string]string{
		"name": "diki",
		"address": "bandung",
	}

	fmt.Printf("%#v\n", employees)
	fmt.Println(employees["name"])
	fmt.Println(employees["address"])
	fmt.Printf("length of employees is %d\n", len(employees))
}