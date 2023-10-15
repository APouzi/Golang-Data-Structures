package patterns

func HappyNumber(n int) bool {
	SquareSumDigit := func(num int) int {
		currsum := 0
		for num > 0 {
			digit := num % 10
			currsum += digit
			num /= 10
		}
		return currsum

	}

	var fast, slow int = n, n

	for {
		slow = SquareSumDigit(slow)
		fast = SquareSumDigit(SquareSumDigit(fast))

		if fast == slow {
			break
		}
	}

	return slow == 1
}