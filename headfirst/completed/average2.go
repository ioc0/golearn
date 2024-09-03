package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(err error) {
	log.Fatal(err)
}
func main() {
	arguments := os.Args[1:]
	var sum float64 = 0
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		check(err)
		sum += number
	}
	fmt.Println(sum)
	fmt.Printf("Average: %0.2f \n", sum/float64(len(arguments)))

}
