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

	countSave := countSaveLists(lists)

	fmt.Println("saveCount:")
	fmt.Println(countSave)
	return nil
}

func countSaveLists(lists [][]int) int {
	saveCount := 0
	for _, list := range lists {
		isSave := isListSave(list)
		if isSave {
			saveCount++
		}
	}
	return saveCount
}

const ASC = 1
const DESC = 2
const UNKNOWN = 0

func isListSave(list []int) bool {
	var previousValuePtr *int
	previousOrder := UNKNOWN

	for _, currentValue := range list {
		if previousValuePtr == nil {
			previousValuePtr = &currentValue
			continue
		}

		currentOrder := UNKNOWN
		if *previousValuePtr < currentValue {
			currentOrder = ASC
		} else if *previousValuePtr > currentValue {
			currentOrder = DESC
		}

		// UNSAFE: previousValue == currentValue
		if currentOrder == UNKNOWN {
			return false
		}

		// UNSAFE: Ordering mismatch
		if previousOrder != currentOrder {
			return false
		}

		previousOrder = currentOrder
		previousValuePtr = &currentValue
	}

	return false
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
