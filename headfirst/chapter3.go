// chapter3
package main

import (
	"fmt"
	"log"
	"os"
)

type point struct {
	x, y int
}

func main3() {
	amount, err := areaCalc(4.2, 3.0)
	fmt.Println(err)
	fmt.Printf("%0.2f needed \n", amount)
}
func fileInfo() {
	fmt.Println("chapter3")
	fileInfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
	//fmt.Printf("the %s cost %d cents each.\n", "gumballs", 23)
	//stringFormat()
	//pseudoChartTable() //p := point{1,2}

}
func stringFormat() {
	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 23)
	fmt.Printf("A string: %s\n", "string")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("Values: %v, %v, %#v, %s", 1.2, "\t", true, "\n")
}
func pseudoChartTable() {
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("______________________________")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 10)
	fmt.Printf("%12s | %2d\n", "Tape", 99)

}
func areaCalc(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10, nil

}
