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
