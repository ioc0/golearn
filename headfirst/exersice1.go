package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main2() {
	guess := rand.Intn(100) + 1
	fmt.Println(guess)
	tries := 10

	for x := 0; x < 10; x++ {
		var suffics string
		if x == 1 {
			suffics = "y"
		} else {
			suffics = "ies"
		}
		nameTry := strings.TrimSpace("tr" + suffics)
		fmt.Println("You have ", tries-x, " ", nameTry)

		fmt.Print("Guess number 1-100: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		answer, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if answer == guess {
			fmt.Println("passed")
			break
		} else if answer > guess {
			fmt.Println("try colder")

		} else {
			fmt.Println("try hotter")

		}

	}

}

// get random func
