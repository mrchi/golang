// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file.
func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64

	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}

	// 按行读取
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}

	// 关闭文件
	err = file.Close()
	if err != nil {
		return numbers, err
	}

	// 扫描文件时出现错误
	if scanner.Err() != nil {
		return numbers, err
	}

	return numbers, nil
}
