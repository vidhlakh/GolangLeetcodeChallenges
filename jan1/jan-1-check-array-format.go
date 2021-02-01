package jan1

// Check Array Formation Through Concatenation
// Can reorder elements of the pieces array
// Can't change order of elements within the piece
// use map to egt O(1)
import (
	"fmt"
	"reflect"
)

func canFormArray(arr []int, pieces [][]int) bool {
	output := false

	pieceMap := make(map[int][]int)
	for _, piece := range pieces {

		pieceMap[piece[0]] = piece
		//piecesMap = append(piecesMap, pieceMap)
	}

	outputArray := make([]int, 0)
	for _, ele := range arr {
		if val, ok := pieceMap[ele]; ok {
			fmt.Println("Valu", val)
			for _, v := range val {
				outputArray = append(outputArray, v)
			}

		}
	}
	fmt.Println("output array", outputArray)
	if reflect.DeepEqual(arr, outputArray) {
		output = true
	}

	return output
}

func main() {

	fmt.Println("Hello Jan 1 ")
	arr := []int{91, 4, 64, 78}
	pieces := [][]int{{78}, {4, 64}, {91}}
	output := canFormArray(arr, pieces)
	fmt.Println("OUTPUT:", output)
}
