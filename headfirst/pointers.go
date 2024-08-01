package main

import (
	"fmt"
)

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}
func double(number *int) { //получаем указатель
	*number *= 2
}
func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}
func mai() {
	thruth := true
	negate(&thruth)
	fmt.Println(thruth)
	var myBool = true
	printPointer(&myBool)
	amount := 6
	double((&amount)) // передали указатель на переменную

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

	fmt.Println(amount)
}
