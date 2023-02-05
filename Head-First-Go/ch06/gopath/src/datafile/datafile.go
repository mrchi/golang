// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file.
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64

	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	// 按行读取
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	// 关闭文件
	err = file.Close()
	if err != nil {
		return nil, err
	}

	// 扫描文件时出现错误
	if scanner.Err() != nil {
		return nil, err
	}

	return numbers, nil
}
