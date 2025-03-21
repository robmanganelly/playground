package main

import "fmt"

type person struct {
	name  string
	age   int
	color string
}

var Fred person

var pet = struct {
	name string
	kind string
}{"Fido", "dog"}

func main() {
	Bill := person{}
	Norma := person{"Norma", 25, "blue"}
	Carol := person{name: "Carol", age: 30}
	pet := 213

	fmt.Println(Fred)
	fmt.Println(Bill)
	fmt.Println(Norma)
	fmt.Printf("hello %s\n", Carol.name)
	fmt.Printf("hello %s, congratulations on your %dth birthday!\n", Carol.name, Carol.age)
	fmt.Println("my pet is", pet)
}
