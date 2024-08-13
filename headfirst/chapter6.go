package main

import (
	"fmt"
	"log"

	"example.com/greetings/datafile"
	"example.com/greetings/ioc0"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
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
	//datafile.GetFloats("data.txt")
	//readfile.Readfile("data.txt")
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3, s4)
	s4[0] = "XXX"
	fmt.Println(s1, s2, s3, s4)
	a1 := []string{"a1", "a2"}
	a1 = append(a1, "a2", "a2")
	a1 = append(a1, "a3", "a3")
	a1 = append(a1, "a4", "a4")
	fmt.Println(a1)
	///////////////////////////
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 6)
	fmt.Println(floatSlice[2], boolSlice[3])
	var intSlice []int
	fmt.Printf("intSlice: %#v \n", intSlice)
	intSlice = append(intSlice, 27)

	intSlice = append(intSlice, 28)
	fmt.Printf("intSlice: %#v \n", intSlice)
	ioc0.YAHU()
	nums, err := datafile.GetFloats("data1.txt")
	check(err)
	fmt.Println(average(nums))

}
func average(slice []float64) float64 {
	num := 0.0
	for _, nums := range slice {
		num += nums
	}
	result := num / float64(len(slice))
	return float64(result)
}
