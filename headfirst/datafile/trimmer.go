package datafile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Trimmer(inputFilename string, symbol string, outputFilename string) {
	fmt.Println(inputFilename)
	fmt.Println(symbol)
	fmt.Println(outputFilename)
	data, _ := trimStrings(inputFilename, symbol)
	writeFile(data, outputFilename)
}

func writeFile(somedata []string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, line := range somedata {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func trimStrings(filename string, symbol string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strings.Trim(line, symbol)
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}
