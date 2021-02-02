package main

import (
	"fmt"
	"strconv"
)

func hammingWeight(num uint32) int {
	var result int
	str := strconv.FormatUint(uint64(num), 2)
	for i := 0; i < len(str); i++ {
		fmt.Println("S:", string(str[i]))
		if string(str[i]) == "1" {
			result++
		}
	}
	return 1
}

func main() {
	fmt.Println("Feb 1 Number of 1- bits")
	var input uint32 = 11
	output := hammingWeight(input)
	fmt.Printf("Number of 1-bits in %d is %d", input, output)
}
