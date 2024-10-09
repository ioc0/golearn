package main

import (
	"fmt"
	"log"

	"example.com/greetings/datafile"
)

func main() {
	fmt.Println("test")
	somenum, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(somenum)

}
