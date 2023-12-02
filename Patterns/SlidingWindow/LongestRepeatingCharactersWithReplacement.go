package slidingwindow

func characterReplacement(s string, k int) int {
	var l int = 0
	var maxFreq, ans int = 0, 0
	var bank map[byte]int = make(map[byte]int)
	for r := 0; r < len(s); r++ {
		bank[s[r]]++
		
		maxFreq = max(maxFreq, bank[s[r]])
		for (r-l+1)-maxFreq > k {
			bank[s[l]]--
			l++
		}
		ans = max(ans, r-l+1)
		if ans > len(s)/2 {
			break
		}
	}
	return ans
}