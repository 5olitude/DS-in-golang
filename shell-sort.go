package main

import "fmt"

func main() {

	arr := []int{9, 6, 5, 8, 0, 7, 4, 3, 1, 2}
	var length = len(arr)
	for gap := length / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < len(arr); i++ {

			j := i

			for {

				if j-gap < 0 || arr[j] >= arr[j-gap] {

					break

				}

				arr[j], arr[j-gap] = arr[j-gap], arr[j]
				j = j - gap

			}

		}

	}
	fmt.Println(arr)
}
