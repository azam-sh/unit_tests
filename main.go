package main

func IndexOfMaxAndMaxElement(numbers []int) (int, int) {
	Index := 0
	maxElement := 0
	for index, value := range numbers {
		if maxElement < value {
			maxElement = value
			Index = index
		}
	}
	return Index, maxElement
}

func IndexOfMaxElement(numbers []int) int {
	Index := 0
	maxElement := numbers[0]
	for index, value := range numbers {
		if maxElement < value {
			maxElement = value
			Index = index
		}
	}
	return Index
}

func main() {
	IndexOfMaxAndMaxElement([]int{10, -5, 1, -100, 25})
}
