package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

import "sorter/algorithm"

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithmName *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithmName)
	}
	values, err := readValues(*infile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Read Values:", values)
	t1 := time.Now()
	switch *algorithmName {
	case "quicksort":
		algorithm.QuickSort(values)
	case "bubblesort":
		algorithm.BubbleSort(values)
	}
	t2 := time.Now()
	fmt.Println("排序程序耗时:", t2.Sub(t1))
	writeValues(values, *outfile)
}

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("打开infile失败:", infile)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		str := string(line)              // 把字节数组转为字符串
		value, err1 := strconv.Atoi(str) // 字符串 转 整数
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create outfile", outfile)
		return err
	}
	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

