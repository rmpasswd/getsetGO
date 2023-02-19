package main

import "fmt"

func main(){
	fmt.Println("Enter Name and Age: ")
	var age int
	var name string
	fmt.Scanf("%s||%d", &name, &age)
	fmt.Printf("Hello %d year old Mr. %s", age, name)
}