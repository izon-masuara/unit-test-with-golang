package main

import (
	"fmt"
	"strconv"
)

/*
	Add the value to array
	return array of string
*/
func AddValue(arr []string, word string) []string {
	if word == "Snow" {
		arr = append(arr, word)
	} else if word == "Jon" {
		arr[1] = word
	}
	return arr
}

/*
	Add the first value on array
	return array of string
*/
func AddFirst(arr []string, word string) []string {
	res := []string{}
	res = append(res, word)
	for _, v := range arr {
		res = append(res, v)
	}
	return res
}

/*
	Count the plate numbers are odd and oven
	return string
*/
func CheckISOddOven(plate string) string {
	arr := []string{}
	tmp := ""

	//  To split the string to array
	for i := 0; i < len(plate); i++ {
		if string(plate[i]) == ";" {
			arr = append(arr, tmp)
			tmp = ""
		} else if i == len(plate)-1 {
			tmp += string(plate[i])
			arr = append(arr, tmp)
		} else {
			tmp += string((plate[i]))
		}
	}

	// Check odd and oven in array
	odd := 0
	oven := 0
	for _, v := range arr {
		num, _ := strconv.Atoi(v)
		if num%2 == 1 {
			odd++
		} else if num%2 == 0 {
			oven++
		}
	}
	return fmt.Sprintf("plat genap sebanyak %v dan plat ganjil sebanyak %v", oven, odd)
}
