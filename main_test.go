package main

import (
	"testing"
)

func testingArr(arrCompare []string, res []string) bool {
	flag := true
	for i, item := range res {
		if arrCompare[i] != item {
			flag = false
		}
	}
	return flag
}

func TestVocal(t *testing.T) {
	// Test 1 add normaly value on array
	arr1 := [...]string{"Rhaegar"}
	arrCompare1 := []string{"Rhaegar", "Snow"}
	res1 := AddValue(arr1[:], "Snow")

	// Test 2 add value at the first index on array
	arr2 := []string{"Snow"}
	arrCompare2 := []string{"jon", "Snow"}
	res2 := AddFirst(arr2, "jon")

	// Test 3 checking odd or oven on the number of plate
	arr3 := "2341;3429;864;1309;1276"
	arrCompare3 := "plat genap sebanyak 2 dan plat ganjil sebanyak 3"

	if testingArr(arrCompare1, res1) != true {
		t.Error("failed test 1")
	}
	if testingArr(arrCompare2, res2) != true {
		t.Error("failed test 2")
	}
	if CheckISOddOven(arr3) != arrCompare3 {
		t.Error("failed test 3")
	}
}
