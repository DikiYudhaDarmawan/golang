//package main

import "fmt"

type Address struct {
	City string
	Province string
	Country string
}

//func main(){
//	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
//	address2 := address1

//	fmt.Println(address1)
//	fmt.Println(address2)

//}

// ampersand
// func main(){
// 	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
// 	address2 := address1

// 	address2.City = "jakarta"

// 	fmt.Println(address1)
// 	fmt.Println(address2)
// }

//asterisk
func main() {
	address1 := Address{"Bsndung" , "Jawa Barat" , "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)
}