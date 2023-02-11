package datafile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening file", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64

	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	// 将defer关键字放在任何普通函数或方法调用之前
	// Go将延迟（也就是推迟）执行函数调用，直到当前函数退出之后。
	// defer 关键字确保函数调用发生，即使调用函数提前退出了。
	defer CloseFile(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
