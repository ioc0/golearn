package main

import (
	"fmt"

	"example.com/greetings/datafile"
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
	fmt.Println(inRange(-100, 105, -12.5, 3.2, 0, 50, 103.5))
	somedat, _ := datafile.GetFloats("test_data.txt")
	average(somedat...)
	average(-100, 105, -12.5, 3.2, 0, 50, 103.5)

}
func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}
func average(array ...float64) {
	sum := 0.0
	for _, num := range array {
		sum += float64(num)
	}
	fmt.Printf("average is : %0.2f \n", (sum / float64(len(array))))
}
