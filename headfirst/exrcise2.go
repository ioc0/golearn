package main

import (
	"fmt"
	"log"
	"math"
)

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("cant get negative root")
	}
	return math.Sqrt(number), nil
}
func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("negative width")
	}
	if height < 0 {
		return 0, fmt.Errorf("negative height")
	}
	area := width * height
	return area / 10.0, nil
}
func divideSomeNum(divident float64, divider float64) (float64, error) {
	if divider == 0 {
		return 0, fmt.Errorf("divider can't be zero")
	}
	return divident / divider, nil
}
func main5() {
	amount, err := paintNeeded(3.22, 4.21)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("you need %.2f litres of paint \n", amount)
	root, err := squareRoot(22.32)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("SQR = %.2f\n", root)
	result, err := divideSomeNum(3.2, 10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("result = %.3f \n", result)
}
