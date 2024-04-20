# My DS Repository in GO

Welcome to my repository where I've shared some useful programs. Below you'll find a brief description of each program along with its corresponding image link.

## 1. Maximum Value Program
![Maximum Value](https://github.com/5olitude/DS-in-golang/blob/25c34d61562494f06a036471802513f6caa4f3e6/images/max.png)

Description: This program calculates the maximum value from a given list of numbers. It's useful for finding the largest element in a dataset.
```go
package main

import "fmt"

func main() {
	numbers := []int{60, 50, 95, 80, 70}
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			temp := numbers[i]
			numbers[i] = numbers[i+1]
			numbers[i+1] = temp

		}

	}
	fmt.Println(numbers)
}
```
[Click here to view the Maximum Value program file](https://github.com/5olitude/DS-in-golang/blob/25c34d61562494f06a036471802513f6caa4f3e6/maximum.go)

## 2. Bubble Sorting Algorithm Program
![Bubble Sorting Algorithm](link_to_bubble_sorting_image)

Description: This program implements the bubble sort algorithm, a simple sorting algorithm that repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order.

[Click here to view the Bubble Sorting Algorithm program file](link_to_bubble_sorting_program_file)

## 3. Minimum Value Program
![Minimum Value](link_to_minimum_value_image)

Description: This program calculates the minimum value from a given list of numbers. It's useful for finding the smallest element in a dataset.

[Click here to view the Minimum Value program file](link_to_minimum_value_program_file)
