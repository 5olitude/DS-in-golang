#  Data Structure and algorithm Repository in Golang

Welcome to my repository where I've shared some useful programs. Below you'll find a brief description of each program along with its corresponding image link.
## Table of Contents
1. [Maximum Value Program](#1-maximum-value-program)
2. [Bubble Sorting Algorithm Program](#2-Bubble-Sorting-Algorithm-Program)
3. [Minimum Value Program](#3-Minimum-Value-Program)
4. [ Selection Sort Program](#4-Selection-Sort-Program)
5. [Linear Table Append](#5-Linear-Table-Append)
6. [Linear Table Insert](#6-Linear-Table-Insert)
7. [Linear Table Delete](#7-Linear-Table-Delete)
8. [Insertion Sort Algorithm](#8-Insertion-Sort-Algorithm)
   
## 1. Maximum Value Program
![Maximum Value](https://github.com/5olitude/DS-in-golang/blob/64b1fccb97dc9be8c3bbf9b882950200e5d70334/images/max.png)

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
package main
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
package main
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
[Click here to view the Minimum Value program file](https://github.com/5olitude/DS-in-golang/blob/44c0643f099386b204e9bdb4d434ff94043ff0ab/min_value.go)

## 4. Selection Sort Program
![Selection Sort](https://github.com/5olitude/DS-in-golang/blob/44c0643f099386b204e9bdb4d434ff94043ff0ab/images/selection.png)

Description: Selection Sort in golang.
```go
package main

import "fmt"

func main() {
	arr := []int{90, 70, 50, 80, 60, 85}

	var min_index int
	for i := 0; i < len(arr)-1; i++ {
		min_index = i
		min_value := arr[min_index]
		for j := i; j < len(arr)-1; j++ {
			if min_value > arr[j+1] {

				min_index = j + 1
				min_value = arr[j+1]
			}

		}
		if i != min_index {
			temp := arr[i]
			arr[i] = arr[min_index]
			arr[min_index] = temp

		}
	}
	fmt.Println(arr)
}
```
[Click here to view the Selection sort program file](https://github.com/5olitude/DS-in-golang/blob/44c0643f099386b204e9bdb4d434ff94043ff0ab/selection-sort.go)

## 5. Linear Table Append 
![Linear Table Append](https://github.com/5olitude/DS-in-golang/blob/554385a8dc491c580c50be068c37a634534de023/images/tableappend.png)
Description: Linear Table Append in Golang.
```go
package main

import "fmt"

func main() {
	arr := []int{90, 70, 50, 80, 60, 85}
	length := len(arr)
	temp := make([]int, length+1)
	for i := 0; i < len(arr); i++ {
		temp[i] = arr[i]

	}
	temp[length] = 75

	fmt.Println(temp)

}

```
[Click here to view the Linear Table append program file](https://github.com/5olitude/DS-in-golang/blob/83cc39fd015be40443d9dbce3d68d443be3b7c64/linear_table_append.go)
## 6. Linear Table Insert
![Linear Table Insert](https://github.com/5olitude/DS-in-golang/blob/d7ffb7138a2ba4c9f7f7863e7146cf199a49536d/images/linear_table_insert.png)
Description: Linear Table Insert in Golang.
```go
package main

import "fmt"

func main() {

	arr := []int{90, 70, 50, 80, 60, 85}
	length := len(arr)
	index := 2
	number := 75
	temp := make([]int, length+1)
	for i := 0; i < length; i++ {

		if i < index {
			temp[i] = arr[i]

		} else {

			temp[i+1] = arr[i]

		}

		temp[index] = number

	}

	fmt.Println(temp)

}



```
[Click here to view the Linear Table insert program file](https://github.com/5olitude/DS-in-golang/blob/70cef3fc1d6f5375d56c877916388dcfc7882d39/linear_table_insert.go)

## 7. Linear Table Delete
![Linear Table Delete](https://github.com/5olitude/DS-in-golang/blob/71ffd491b14743fdcc644662c9aac114fb2d02cb/images/linear_table_delete.png)
Description: Linear Table Delete in Golang.
```go
package main

import "fmt"

func main() {
	arr := []int{90, 70, 50, 80, 60, 85}
	length := len(arr)
	temp := make([]int, length-1)
	index := 2

	for i := 0; i < length; i++ {
		if i < index {
			temp[i] = arr[i]

		}
		if i > index {
			temp[i-1] = arr[i]

		}

	}
	fmt.Println(temp)
}

```
[Click here to view the Linear Table delete program file](https://github.com/5olitude/DS-in-golang/blob/351623794d6ffe39d1d0950b4be7a0d543c258e9/linear_table_delete.go)
## 8. Insertion Sort Algorithm
![Insertion Sort Algorithm](https://github.com/5olitude/DS-in-golang/blob/c28efa6b7a650d154eec94243c10d67e470117a9/images/insert.png)
Description: Insertion Sort  in Golang.
```go
package main

import "fmt"

func main() {
	arr := []int{90, 70, 50, 80, 60, 85}
	for i := 0; i < len(arr); i++ {
		insert_element := arr[i]
		insert_position := i
		for j := insert_position - 1; j >= 0; j-- {
			if arr[j] > insert_element {

				arr[j+1] = arr[j]
				insert_position--

			}

			arr[insert_position] = insert_element
		}

	}
	fmt.Println(arr)
}


```
[Click here to view the Insertion Sort  program file](https://github.com/5olitude/DS-in-golang/blob/e96c4c712c2b907e5316fc69f0c4c9902346689d/insert_sort.go)
