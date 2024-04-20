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
![Bubble Sorting Algorithm](https://github.com/5olitude/DS-in-golang/blob/25c34d61562494f06a036471802513f6caa4f3e6/images/bubble.png)

Description: This program implements the bubble sort algorithm, a simple sorting algorithm that repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order.
```go
import "fmt"

func main() {
	numbers := []int{60, 50, 95, 80, 70}
	for i := 0; i < len(numbers)-1; i++ {

		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				temp := numbers[j+1]
				numbers[j+1] = numbers[j]
				numbers[j] = temp

			}
		}

	}
	fmt.Println(numbers)
}
```

[Click here to view the Bubble Sorting Algorithm program file](https://github.com/5olitude/DS-in-golang/blob/1aad678b0224b1dc60238f8c1344ae6d95936525/bubble_sort.go)

## 3. Minimum Value Program
![Minimum Value](https://github.com/5olitude/DS-in-golang/blob/25c34d61562494f06a036471802513f6caa4f3e6/images/min.png)

Description: This program calculates the minimum value from a given list of numbers. It's useful for finding the smallest element in a dataset.
```go
import "fmt"

func main() {
	numbers := []int{60, 80, 95, 50, 70}
	min_index := 0
	for j := 1; j < len(numbers); j++ {
		if numbers[min_index] > numbers[j] {
			min_index = j

		}

	}
	fmt.Println(numbers[min_index])
}
```
[Click here to view the Minimum Value program file](https://github.com/5olitude/DS-in-golang/blob/25c34d61562494f06a036471802513f6caa4f3e6/images/min.png)
