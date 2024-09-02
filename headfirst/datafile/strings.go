package datafile

import (
	"bufio"
	"os"
)

func check(err error) error {
	return err
}
func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	check(err)
	check(scanner.Err())
	return lines, err

}
