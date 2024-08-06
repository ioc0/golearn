// arrays
package main

import (
	"fmt"
)

func main() {
	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	primes := [5]int{1, 2, 3, 4, 5}
	fmt.Println(notes[3])
	fmt.Println(primes[3])
	fmt.Println(notes)
	for i := 0; i <= len(notes)-1; i++ {
		fmt.Println(i, notes[i])
	}
	for _, prime := range primes {
		fmt.Println(prime)
	}
	for index, _ := range primes {
		fmt.Println(index)

	}

	var summMeat float64
	meat := [3]float64{79.2, 56.4, 89.1}
	for _, part := range meat {
		summMeat += part
	}
	result := summMeat / float64(len(meat))
	fmt.Printf("average meat needed %.2f", result)

}
