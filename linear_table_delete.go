package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 9, 4}
	length := len(arr)
	temp := make([]int, length-1)
	index := 0

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

