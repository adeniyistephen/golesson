package main //<== main is the starting point of the application.

import "fmt"

func main() {
	fmt.Println("Hello World") //<== to print out in the console.

	//setting of variables in golang
	var string = "stephen"
	var int = 12
	var float = 0.5
	var boolean = false
	// you can only use this type of variable declaration inside of a function.
	name := "oluwatola"
	age := 22
	decimal := 1.0
	trueOrfales := false
	firstname, lastname := "stephen", "adeniyi"
	fmt.Println(firstname, lastname)
	fmt.Printf("%T, %T, %T, %T, %T, %T, %T, %T", string,int,float,boolean,name,age,decimal,trueOrfales) //<== to print out variable types in golang other's can be checked on the fmt golang doc
} 