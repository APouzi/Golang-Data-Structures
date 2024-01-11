package arrays

// Given a string s, return the length of the longest substring between two equal characters, excluding the two characters. If there is no such substring return -1.
// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s = "aa"
// Output: 0
// Explanation: The optimal substring here is an empty substring between the two 'a's.

// Example 2:
// Input: s = "abca"
// Output: 2
// Explanation: The optimal substring here is "bc".

// Example 3:
// Input: s = "cbzxy"
// Output: -1
// Explanation: There are no characters that appear twice in s.

//Other TestCases: "mgntdygtxrvxjnwksqhxuxtrv" = 18, "ojdncpvyneq" = 4, "scayofdzca" = 6

func maxLengthBetweenEqualCharacters(s string) int {
	var bank map[byte][]int = make(map[byte][]int)
	for i :=0;i<len(s);i++{
		bank[s[i]] = append(bank[s[i]], i)
	}
	var ans int = 0
	for _, v := range bank{
		if len(v) > 1{
			ans = max(ans, (v[len(v)-1] - v[0])-1)
		}
	}
	return ans
}