package jan6

import "fmt"

// FindKthPositive find kth missing positve number
func FindKthPositive(arr []int, k int) int {
	result := make([]int, 0)
	counter := 1
	i := 0
	for i < len(arr) {

		if counter != arr[i] {
			result = append(result, counter)

		} else if counter == arr[i] {
			i++
		}

		counter++

	}

	for {
		if len(result) >= k {
			break
		}

		result = append(result, counter)
		counter++

	}

	return result[k-1]
}
func main() {
	input := []int{2, 3, 4, 7, 11}
	output := FindKthPositive(input, 9)
	fmt.Println("OUTPUT", output)
}
