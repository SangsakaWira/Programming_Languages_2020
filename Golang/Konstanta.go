package main

import "fmt"

func main(){
	const firstName string = "john"
	fmt.Print("halo ", firstName, "!\n")
	const value = (((2 + 6) % 3) * 4 - 2) / 3
	fmt.Println(value)
	
}

// go run Konstanta.go