package main

import (
	"fmt"
	"keyboard"
	"log"
)

func grade() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >=60{
		status ="passing"
	}else{
		status = "failing"
	}
	fmt.Println("A grade of",grade,"is",status)
}
func fahrenheit(){
	fmt.Print("temp in f: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n",celsius)


}
func main1(){
	grade()
	fahrenheit()
}
