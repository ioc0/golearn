package main

import (
	"fmt"
	"log"

	"example.com/greetings/datafile"
)

func main() {
	count("votes.txt")
	lines, err := datafile.ReadStrings("votes.txt")
	if err != nil {
		log.Fatalln(err)
	}
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}

	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}

func count(filename string) {
	lines, err := datafile.ReadStrings(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
