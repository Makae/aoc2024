package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

func main() {
	err := runMain()
	if err != nil {
		fmt.Errorf(err.Error())
	}
}

func runMain() error {
	list1, list2, err := readLists()
	if err != nil {
		return err
	}

	sort.Ints(list1)
	sort.Ints(list2)

	diffSum := 0
	diffSum = getDiffSum(list1, diffSum, list2)

	similaritySum := 0
	similaritySum = getSimilaritySum(list1, list2, similaritySum)

	fmt.Println("Diff:", diffSum)
	fmt.Println("Similarity:", similaritySum)
	return nil
}

func getDiffSum(list1 []int, diffSum int, list2 []int) int {
	for i := 0; i < len(list1); i++ {
		diffSum += getDiff(list1[i], list2[i])
	}
	return diffSum
}

func getSimilaritySum(list1 []int, list2 []int, similaritySum int) int {
	count1 := getNumberCount(list1)
	count2 := getNumberCount(list2)
	for key, count1Value := range count1 {
		count2Value, ok := count2[key]
		if !ok {
			count2Value = 0
		}
		similaritySum += count1Value * count2Value
	}
	return similaritySum
}

func getDiff(i1 int, i2 int) int {
	if i1 < i2 {
		return i2 - i1
	}

	return i1 - i2
}

func getNumberCount(list []int) map[int]int {
	result := make(map[int]int)
	for i := 0; i < len(list); i++ {
		existing, ok := result[i]
		if !ok {
			result[i] = 0
			existing = 0
		}

		result[i] = existing + 1
	}
	return result
}

func readLists() ([]int, []int, error) {
	reader := bytes.NewBuffer(input)
	scanner := bufio.NewScanner(reader)

	list1 := make([]int, 0)
	list2 := make([]int, 0)
	for scanner.Scan() {

		splitData := strings.Split(scanner.Text(), "   ")
		value1, err := strconv.Atoi(splitData[0])
		if err != nil {
			return nil, nil, err
		}
		value2, err := strconv.Atoi(splitData[1])
		if err != nil {
			return nil, nil, err
		}
		list1 = append(list1, value1)
		list2 = append(list2, value2)
	}
	return list1, list2, nil
}
