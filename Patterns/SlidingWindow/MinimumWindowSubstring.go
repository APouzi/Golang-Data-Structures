package slidingwindow

// Given two strings s and t of lengths m and n respectively, return the minimum window
// substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".
// The testcases will be generated such that the answer is unique.

// Example 1:
// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

// Example 2:
// Input: s = "a", t = "a"
// Output: "a"
// Explanation: The entire string s is the minimum window.

// Example 3:
// Input: s = "a", t = "aa"
// Output: ""
// Explanation: Both 'a's from t must be included in the window.
// Since the largest window of s only has one 'a', return empty string.


func MinWindow(s string, t string) string {
	var compareTo map[byte]int = make(map[byte]int)
	var window map[byte]int = make(map[byte]int)
	for i := 0; i < len(t); i++ {
		compareTo[t[i]]++
	}

	var l int = 0
	currHave, need := 0, len(compareTo) // we have the need be the size of the array due to the fact that we will only grow currHave if the value of given character like compareto[s[r]] == bank[s[r]]. That way they match like that. 
	res := [2]int{-1, -1}
	var reslen int = 777777777
	for r := 0; r < len(s); r++ {
		window[s[r]]++
		if window[s[r]] == compareTo[s[r]] {//This grows currHave ONLY if we have the same given variables at the character keys, otherwise, there is no need to compare the two.
			currHave++
		}
		for currHave == need {
			if (r - l + 1) < reslen {//Since we are in the zone where we have the minimum characters, we will continue to decrease the window. 
				res[0], res[1] = l, r
				reslen = (r - l + 1)
			}
			window[s[l]]--
			if window[s[l]] < compareTo[s[l]] { //This window is important because we only decrease when its the same character as compare too that has been removed and also importantly, the window's key has less repeatition than the compare too. This is because sometimes, we will have compareTo{a:1,b:1,c:1} and bank{a:1,b:2,c:1}, we will have to remove B at least twice before this condition is met. (look s = "ADOBECODEBANC", t = "ABC")
				currHave--
			}
			l++
		}
	}
	if reslen != 777777777 { 
		return s[res[0] : res[1]+1]
	}
	return ""//If result never changed because there is no comparision, then return an empty string.
}