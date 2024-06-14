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
