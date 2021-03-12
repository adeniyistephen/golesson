package main //<== main is the starting point of the application.

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/adeniyistephen/golesson/pkg"
)

func main() {
	fmt.Println("Hello World") //<== to print out in the console.

	//setting of variables in golang
	var stringVar = "stephen"
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
	fmt.Printf("%T, %T, %T, %T, %T, %T, %T, %T", stringVar, int, float, boolean, name, age, decimal, trueOrfales) //<== to print out variable types in golang other's can be checked on the fmt golang doc

	//we have different types of packages in golang, eg the "math" package.
	fmt.Println("\n", math.Ceil(2.7))

	//importing of packages in golang...
	pkg.Pkg()

	//executing functions in golang
	fmt.Println("\n", sum(4, 5))

	//arrays and slice in golang
	fruitArr := []string{"apple", "orange", "pine-apple", "strawberry"}      // declare a slice
	fruitArr = append(fruitArr, "banana")                                    //append to a slice
	fmt.Printf("len=%d cap=%d %v\n", len(fruitArr), cap(fruitArr), fruitArr) // check for the lenth of a slice and capacity
	fmt.Println(fruitArr[1:4])

	//loops with condition
	countFizz := 0
	countBuzz := 0
	countFizzBuzz := 0
	rangeCar := []string{"benz", "toyota", "mitsibushi", "rover", "honda"}
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			countFizz++
			fmt.Println("FizzBuzz = ", countFizz)
		} else if i%3 == 0 {
			countBuzz++
			fmt.Println("Fizz = ", countBuzz)
		} else if i%5 == 0 {
			countFizzBuzz++
			fmt.Println("Buzz = ", countFizzBuzz)
		} else {
			fmt.Println(i)
		}
	}
	// use of range
	for i, car := range rangeCar {
		fmt.Printf("index=%d car=%s\n", i, car)
	}
	//########################################################

	//mapping with some struct in it.
	type boy struct {
		name string
		age  int64
	}

	mapBoy := make(map[int64]boy)
	mapBoy[1] = boy{
		name: "stephen",
		age:  18,
	}
	mapBoy[2] = boy{
		name: "adeniyi",
		age:  18,
	}
	fmt.Println(mapBoy, "\n", mapBoy[1])
	//###############################################

	// Running Interface
	p := person{
		name:        "steve",
		yearBirth:   1990,
		currentYear: 2021,
	}
	fmt.Println("age: ", getAge(p))

	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

//functions in golang
func sum(a int, b int) int { //<== this function has an input parameter of a and b of type integer and returns an integer.
	return a + b
}

//###########################################
//running interface
type inter interface { //declare the interface
	calculateAge() int
}

type person struct { // declare the class that will implement the interface
	name        string
	yearBirth   int
	currentYear int
}

func (p person) calculateAge() int { // impliment the interface
	var age int
	age = p.currentYear - p.yearBirth
	return age
}

func getAge(i inter) int { // call the interface
	return i.calculateAge()
}

//##########################################################

//Info ...
type Info struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func index(w http.ResponseWriter, r *http.Request) {
	var i = Info{
		Name:  "stephen",
		Email: "stephen@gmail.com",
		ID:    "232ere-ewewwe443-wer4r",
	}

	js, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
