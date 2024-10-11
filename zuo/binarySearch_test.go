package zuo

import (
	"math/rand"
	"testing"
)

func generateNums() (left, right, want int) {
	// 随机生成 left 和 right要求是整型 0- 100
	left = rand.Int() % 100
	right = rand.Int() % 100
	want = (left + right) / 2
	return
}

func TestFindIndex(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	var tests []struct {
		name string
		args args
		want int
	}
	for range 100 {
		left, right, want := generateNums()
		tests = append(tests, struct {
			name string
			args args
			want int
		}{name: "TestFindIndex", args: args{left: left, right: right}, want: want})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIndex(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("FindIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintBinary(t *testing.T) {
	a := 78
	PrintBinary(int32(a))
}
