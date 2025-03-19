func isPalindrome(x int) bool {

	reversedX := 0
	remains := x
	for remains > 0 {
		lastDigit := remains % 10              //1
		reversedX = reversedX*10 + lastDigit   // 21
		remains = ((remains - lastDigit) / 10) // 1-1/10 = 0
	}

	// fmt.Println("x ", x, " re ", reversedX)

	if x == reversedX {
		return true
	}
	return false
}
