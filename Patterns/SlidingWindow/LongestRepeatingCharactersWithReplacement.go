package slidingwindow

func CharacterReplacement(s string, k int) int {
	var l int = 0
	var maxFreq, ans int = 0, 0 //MaxFreq will be tracking which character has the highest frequency.
	var bank map[byte]int = make(map[byte]int)
	for r := 0; r < len(s); r++ {
		bank[s[r]]++

		maxFreq = max(maxFreq, bank[s[r]]) // This will keep track of whether the currently added byte into the bank is the most frequent. If it's not, the last byte that was updated, is the answer. This can either be the last one or multiple iterations before.
		for (r-l+1)-maxFreq > k {          //This the "Trick" we use to know whether there is enough replacement points (k) to replace whatever characters to make the string longer.
			bank[s[l]]--
			l++
		}
		ans = max(ans, r-l+1)
		// if ans > len(s)/2 {
		// 	break
		// }
	}
	return ans
}