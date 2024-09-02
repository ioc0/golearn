package main

import (
	"fmt"

	"github.com/headfirstgo/datafile"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
	mySlice2 := myArray2[1:4]
	fmt.Println(mySlice2)
	array1 := [8]string{"a", "b", "c", "d", "e", "f", "j", "h"}
	slice1 := array1[0:3]
	//slice2 := array1[2:5]
	array1[1] = "X"
	fmt.Println(array1, slice1)
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))
	slice = append(slice, "d", "e", "f")

	fmt.Println(slice, len(slice))
	data, _ := datafile.GetFloats("data.txt")
	fmt.Println(data)

}
