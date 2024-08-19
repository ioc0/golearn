package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		check(err)
		numbers = append(numbers, number)
	}
	err = file.Close()
	check(err)
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, err
}
