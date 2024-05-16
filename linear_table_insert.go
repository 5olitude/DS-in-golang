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
