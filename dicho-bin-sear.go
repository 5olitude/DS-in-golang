package main

import "fmt"

func main() {
	search_value := 1
	arr := []int{1, 2, 3, 4, 5, 6} // must be in sorted order
	mid := 0
	low := 0
	high := len(arr) - 1
	for {
		if low >= high {
			break
		}
		mid = (low + high) / 2
		if arr[mid] == search_value {
			fmt.Println(mid)
			return

		} else if arr[mid] > search_value {
			high = mid - 1
		} else if arr[mid] < search_value {
			low = mid + 1

		}

	}
	fmt.Println("element not found ")
}
