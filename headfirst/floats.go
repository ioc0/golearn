package datafile

import (
        "bufio"
	"os"
	"strconv"
)
func check(e error){
	if e != nil {
	panic(e)
}
}

func GetFloats(fileName string)([]float64, error){
	var numbers []float64
	file, err : os.Open(fileName)
	check(err)
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParceFloat(scanner.Text(),64)
		check(err)
		i++
	}
	err = file.Close()
	check(err)
	check(scanner.Err())
	return numbers, nil
}
