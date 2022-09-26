package main

import "fmt"

//Interface
type Animal interface {
	speak()
}

type Dog struct {
}

func (d Dog) speak() {
	fmt.Println("woaw woaw")
}

func (d Dog) move() {
	fmt.Println("4 chân")
}

//multiple interface
type Movement interface {
	move()
}

func main() {
	// var animal Animal

	// animal = Dog{}

	// animal.speak()

	dog := Dog{}

	var m Movement = dog
	m.move()
	var a Animal = dog
	a.speak()
}
