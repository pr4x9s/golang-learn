package main

import "fmt"

func main() {
	// var name string = "golang"

	// inferred
	var name = "golang"
	var isAdult = true
	var age = 30

	// shorthand syntax to declare variable and assign value to it
	name1 := "Prantik"

	// declare variable but assign value later
	var abc string
	var price float32

	abc = "hello string"
	price = 50.5

	fmt.Println(name, isAdult, age, name1, abc, price)
}
