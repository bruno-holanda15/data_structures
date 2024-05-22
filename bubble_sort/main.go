package main

import "fmt"

func main() {
	// Swap, check order, one-round bubbling, repeated n times...
	numbers := []int{6, 21, 1, 4, 2, 1, 9, 7, 13}

	for range numbers{
		for i :=0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				swap(numbers, i, i+1)
			}
		}
	}

	fmt.Println(numbers)
}

func swap(numbers []int, i, j int) {
	// (2 1) -> (1 2)
	// (2 1)  temp=2 -> (1 2)
	temp := numbers[i]
	numbers[i] = numbers[j]
	numbers[j] = temp
}
