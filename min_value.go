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
