package main

import (
	"fmt"
)

func main() {
	primes := []int{1, 2, 2, 33, 4, 4, 1, 1}
	primes[0] = 2
	primes[1] = 5
	fmt.Println(primes[5])
	fmt.Println("//////////////////////////////////////////////////")
	notes := []string{"do", "re", "mi", "fa", "sol"}
	fmt.Println(len(notes))
	fmt.Println(len(primes))
	for i, num := range primes {
		fmt.Print(i, " ", num, "\n")

	}
	fmt.Println("//////////////////////////////////////////////////")
	for _, string := range notes {
		fmt.Println(string)
	}
	fmt.Println("//////////////////////////////////////////////////")
	letters := []string{"a", "b", "c"}
	for _, letter := range letters {
		fmt.Println("letter: ", letter)
	}
	fmt.Println("//////////////////////////////////////////////////")
	myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mySlice := myArray[1:8]
	for _, num := range mySlice {
		fmt.Println(num)
	}
	fmt.Println("//////////////////////////////////////////////////")
	myArray2 := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(len(myArray2))

}
