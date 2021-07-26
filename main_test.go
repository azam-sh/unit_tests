package main

import "testing"

func TestIndexOfMaxAndMaxElement(t *testing.T) {
	tests := []struct {
		name           string
		numbers        []int
		IndexWant      int
		MaxElementWant int
	}{
		{"All of numbers are positive", []int{1, 4, 25, 25, 100}, 4, 100},
		//{"All of numbers are negative", []int{-1, -4, -25, -25, -100}, 0, -1},
		{"Positive and negative numbers", []int{-1, 5, -20, 2}, 1, 5},
	}
	for _, test := range tests {
		index, element := IndexOfMaxAndMaxElement(test.numbers)
		if index != test.IndexWant || element != test.MaxElementWant {
			t.Errorf("IndexOfMaxandMaxElement test %s gotIndex %d IndexWant %d gotElement %d ElementWant %d", test.name, index, test.IndexWant, element, test.MaxElementWant)
		}
	}
}
