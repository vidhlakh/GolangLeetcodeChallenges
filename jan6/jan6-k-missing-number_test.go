package jan6

import (
	"fmt"
	"testing"
)

func TestFindKthPositive(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{2, 3, 4, 7, 11}, 5}, 9},
		{"case2", args{[]int{1, 2, 3, 4}, 2}, 6},
		{"case3", args{[]int{}, 7}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthPositive(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("FindKthPositive() = %v, want %v", got, tt.want)
			}
			fmt.Println("WANT", tt.want)
		})
	}
}
