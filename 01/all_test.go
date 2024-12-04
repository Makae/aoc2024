package main

import (
	"fmt"
	"testing"
)

func Test_getNumberCount(t *testing.T) {
	count := getNumberCount([]int{
		1, 1, 1,
		23, 34,
		1,
		12222, 12222, 12222,
		1, 1,
	})
	if c, _ := count[1]; c != 6 {
		t.Errorf("Count of 1 is not 6, is %d", c)
	}

	if c, _ := count[23]; c != 1 {
		t.Errorf("Count of 23 is not 1, is %d", c)
	}

	if c, _ := count[34]; c != 1 {
		t.Errorf("Count of 34 is not 1, is %d", c)
	}

	if c, _ := count[12222]; c != 3 {
		t.Errorf("Count of 12222 is not 3, is %d", c)
	}
}

func Test_Similarity(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	count1 := getNumberCount(list1)
	count2 := getNumberCount(list2)

	fmt.Println(count1, count2)

	similarity := getSimilaritySum(list1, list2)

	if similarity != 31 {
		t.Errorf("Similarity is not 31, is %d", similarity)
	}
}
