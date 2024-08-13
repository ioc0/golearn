package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	arguments := os.Args[1:]
	var sum float64 = 0
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		check(err)
		sum += number
	}
	sampleCount := float64(len(arguments))
	fmt.Printf("Average: %.02f\n", sum/sampleCount)

}
