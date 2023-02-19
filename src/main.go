package main

import "fmt"

func main(){
	var students [3]string;		// conventional way
	students[0] = "MA"
	students[1] = "TA"
	students[2] = "QA"
	
	fmt.Println(len(students))
	fmt.Println(students[2])
	
	teachers := [3]string{"YS", "df"}	// shorthand, string literals 
	teachers[1] = "MAD"
	teachers[2] = "JA"
	fmt.Println(teachers)
	// or use teachers := [...]string{"YS", "df"}


		////////////SLICE///////////
		
}