package arrays

func IsAnagramHashMap(s string, t string) bool {
	var tl, sl int = len(s), len(t)
	if sl != tl {
		return false
	}
	var sm, tm map[byte]int = make(map[byte]int, tl), make(map[byte]int, sl)
	for i := 0; i < tl; i++ {
		sm[s[i]]++
		tm[t[i]]++
	}

	for i := 0; i < tl; i++ {
		if sm[s[i]] != tm[s[i]] {
			return false
		}
	}
	return true
}


