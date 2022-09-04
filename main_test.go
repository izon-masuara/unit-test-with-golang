package main

import "testing"

func TestAdd(t *testing.T) {
	result1 := Add(2, 3)
	result2 := Add(10, 2)
	result3 := Add(23, 21)

	if result1 != 5 {
		t.Error("Failed equal 5")
	} else if result2 != 12 {
		t.Error("Failed equal 12")
	} else if result3 != 44 {
		t.Error("Failed equal 44")
	}
}
