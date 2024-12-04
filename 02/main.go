package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

func main() {
	err := runMain()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
}

func runMain() error {
	lists, err := readInput()
	if err != nil {
		return err
	}

	fmt.Println("lists:")
	fmt.Println(lists)
	return nil
}

func getDiff(i1 int, i2 int) int {
	if i1 < i2 {
		return i2 - i1
	}

	return i1 - i2
}

func readInput() ([][]int, error) {
	reader := bytes.NewBuffer(input)
	scanner := bufio.NewScanner(reader)

	list := make([][]int, 0)
	for scanner.Scan() {

		splitData := strings.Split(scanner.Text(), " ")
		intList := make([]int, len(splitData))
		for idx, elm := range splitData {
			intVal, err := strconv.Atoi(elm)
			if err != nil {
				return nil, err
			}
			intList[idx] = intVal
		}
		list = append(list, intList)
	}
	return list, nil
}
