// You can edit this code!
// Click here and start typing.
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
