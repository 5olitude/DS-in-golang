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
