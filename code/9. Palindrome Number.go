package code

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var res int
	for i := x; i > 0; i /= 10 {
		res = res*10 + i%10
	}
	return res == x
}
