package arrays

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