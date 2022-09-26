package main

import "fmt"

//Interface
type Animal interface {
	speak()
}

type Dog struct {
}

func (d *Dog) speak() {
	fmt.Println("woaw woaw")
}

func (d *Dog) move() {
	fmt.Println("4 chân")
}

//multiple interface
type Movement interface {
	move()
}

//Embed interface
type NextAnimal interface {
	Movement
	Animal
}

//Empty interface
func goout(i interface{}) {
	fmt.Println(i)
}

type data struct {
	index int
}

func main() {
	// var animal Animal

	// animal = Dog{}

	// animal.speak()

	dog := Dog{}

	// var m Movement = dog
	// m.move()
	// var a Animal = dog
	// a.speak()
	var na NextAnimal = &dog
	na.speak()
	na.move()

	goout(10)
	goout(10.15)

	data := data{1002}
	goout(data)
}
