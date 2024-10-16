package code

import "testing"

func TestFindPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 2, 3, 1}}, 2},
		{"2", args{[]int{1, 2, 1, 3, 5, 6, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); got != tt.want {
				t.Errorf("FindPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
