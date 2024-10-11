package zuo

import "fmt"

func FindIndex(left, right int) int {
	return left + (right-left)/2
}

func PrintBinary(n int32) {
	for i := 0; i < 32; i++ {
		fmt.Print((n >> i) & 1)
	}
}
