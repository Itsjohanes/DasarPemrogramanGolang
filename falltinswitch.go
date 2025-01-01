package main

import "fmt"

func main(){
	var point = 8

	switch point {
	case 8:
		fmt.Println("Perfect")
		fallthrough
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}
}
