package slidingwindow

func BuyHighSellLow(prices []int) int {
	lowest := prices[0] //Originally did -999999, but that would mean that we could possibly get a worse conidition
	maxSell := -9999999999

	for _, v := range prices {
		possibleMS := v - lowest
		if lowest > v {
			lowest = v
		}

		if possibleMS > maxSell {
			maxSell = possibleMS
		}
	}
	return maxSell
}