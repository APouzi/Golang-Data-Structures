package slidingwindow

func MinWindow(s string, t string) string {
	var compareTo map[byte]int = make(map[byte]int)
	var window map[byte]int = make(map[byte]int)
	for i := 0; i < len(t); i++ {
		compareTo[t[i]]++
	}

	var l int = 0
	currHave, need := 0, len(compareTo)
	res := [2]int{-1, -1}
	var reslen int = 777777777
	for r := 0; r < len(s); r++ {
		window[s[r]]++
		if _, ok := compareTo[s[r]]; ok && window[s[r]] == compareTo[s[r]] {
			currHave++
		}
		for currHave == need {
			if (r - l + 1) < reslen {
				res[0], res[1] = l, r
				reslen = (r - l + 1)
			}
			window[s[l]]--
			// if window[s[l]] == 0{
			//     delete(window, s[l])
			// }
			if _, ok := compareTo[s[l]]; ok && window[s[l]] < compareTo[s[l]] {
				currHave--
			}
			l++
		}
	}
	if reslen != 777777777 {
		return s[res[0] : res[1]+1]
	}
	return ""
}