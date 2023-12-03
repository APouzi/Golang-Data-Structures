package arrays

//Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// Example 2:
// Input: strs = [""]
// Output: [[""]]

// Example 3:
// Input: strs = ["a"]
// Output: [["a"]]

func GroupAnagrams(strs []string) [][]string {
	var groups map[[26]int][]string = map[[26]int][]string{}
	for i := 0; i < len(strs); i++ {
		var count [26]int
		for c := 0; c < len(strs[i]); c++ {
			count[strs[i][c]-'a']++
		}
		groups[count] = append(groups[count], strs[i])
	}
	ans := [][]string{}
	for _, v := range groups {
		ans = append(ans, v)
	}
	return ans
}