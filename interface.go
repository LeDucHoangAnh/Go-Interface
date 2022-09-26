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

func main() {
	var animal Animal

	animal = Dog{}

	animal.speak()
}
