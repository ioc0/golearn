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
	testMap()
}

func count(filename string) {
	lines, err := datafile.ReadStrings(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
func testMap() {
	ranks := make(map[string]int)
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])
	elements := make(map[string]string)
	elements["Li"] = "Lithium"
	fmt.Println(elements["Li"])
	someMap := map[string]int{"a": 1, "b": 2}
	fmt.Println(someMap["a"])

}
