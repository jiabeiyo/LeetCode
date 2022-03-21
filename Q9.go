func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := x
	z := 0
	for x != 0 {
		z = z*10 + x%10
		x /= 10
	}
	return y == z
}