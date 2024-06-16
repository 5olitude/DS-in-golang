package main

import "fmt"

func main() {
	arr := []int{50, 60, 70, 80, 90}
	middle := len(arr) / 2
	for i := 0; i <= middle; i++ {
		temp := arr[i]
		arr[i] = arr[len(arr)-i-1]
		arr[len(arr)-i-1] = temp

	}
	fmt.Println(arr)
}
