package main

import "fmt"

func main(){
	var firstName string = "john"

    var lastName string
    lastName = "wick"

	fmt.Printf("halo %s %s!\n", firstName, lastName)
	
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	var message string = "Halo"
	fmt.Printf("message: %s \n", message)
}

// go run DataTypes.go