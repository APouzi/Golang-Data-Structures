package slidingwindow

// You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.
// Return the length of the longest substring containing the same letter you can get after performing the above operations.

// Example 1:
// Input: s = "ABAB", k = 2
// Output: 4
// Explanation: Replace the two 'A's with two 'B's or vice versa.

// Example 2:
// Input: s = "AABABBA", k = 1
// Output: 4
// Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.
// There may exists other ways to achieve this answer too.

func CharacterReplacement(s string, k int) int {
	var l int = 0
	var maxFreq, ans int = 0, 0 //MaxFreq will be tracking which character has the highest frequency.
	var bank map[byte]int = make(map[byte]int)
	for r := 0; r < len(s); r++ {
		bank[s[r]]++

		maxFreq = max(maxFreq, bank[s[r]]) // This will keep track of whether the currently added byte into the bank is the most frequent. If it's not, the last byte that was updated, is the answer. This can either be the last one or multiple iterations before.
		for (r-l+1)-maxFreq > k {//This the "Trick" we use to know whether there is enough replacement points (k) to replace whatever characters to make the string longer. The reason this works is because of this, maxFreq is going to be the highest repeated character. This is important because lets say we have "AAAA", meaning our maxFreq would be 4 by the end of the line. Now if we do "3-0+1", which would be 4. So (r-l+1=4) - MaxFreq=4 >k. We are asking this because if that is bigger than k, it means that we need to decrement the left pointer and then increment l++.
			bank[s[l]]--
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}