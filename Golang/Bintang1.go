package main

import "fmt"

func main(){
	var n = 5
	for i:=0; i<n;i++{
		for j:=0; j<n;j++{
			if(j<=i){
				fmt.Print("*")
			}else{
				fmt.Print("o")
			}
		}
		fmt.Print("\n")
	}
}