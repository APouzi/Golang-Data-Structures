package slidingwindow

//Given a string s, find the length of the longest substring without repeating characters.

// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

//testcases: "qrsvbspk" = 5, " " = 1
//first one will fail without a for loop inside deleting instead of if loop.

func LengthOfLongestSubstring(s string) int {
    if s == ""{
        return 0
    }
    bank := map[byte]bool{}
    var l int = 0
    ans := 0
    for r :=0; r <len(s);r++{
        for bank[s[r]]{
            delete(bank,s[l])
            l++
        }
        bank[s[r]] = true
        ans = max(len(bank), ans)
    }
    return ans
}