package arrays

func IsAnagramHashMap(s string, t string) bool {
	var tl, sl int = len(s), len(t)
	if sl != tl {
		return false
	}
	var sm, tm map[byte]int = make(map[byte]int, tl), make(map[byte]int, sl)
	//Essentially what we want to do is add s and t strings to our hashmaps at the same time, since they are of equal length. Increment the count of each letter.
	for i := 0; i < tl; i++ {
		sm[s[i]]++
		tm[t[i]]++
	}
	//Once we went through both strings and added them to the hashmaps, what we want to do is then ask if every single string in one of the strings are there in both hashmaps. Notice that we are only doing one string at a time and thats it.
	for i := 0; i < tl; i++ {
		if sm[s[i]] != tm[s[i]] {
			return false
		}
	}
	return true
}

func IsAnagramStringTrick(s string, t string) bool {
	var tl, sl int = len(s), len(t)
	if sl != tl {
		return false
	}
	//Here we are going to be creating an array that is the size of the english alphabet and that will be 26 letters. The reason for this is that we are going to be checking every single letter in the two given strings and this array will allow us to do this using the ASCII letters.
	var sCheck []int = make([]int, 26)
	//Using ASCII, we realize that the letters are lowercase and this means that "a" starts at 97 and b is "98" and this continues for the rest of the lowercase alphabet. So what happens here is that say we have the letter "a" from string "t", now when we subtract the letter "a", which is ASCII, we are in fact subtracting 97 from 97 and that would be 0. So if we do "t[i]-97" it means that we are getting the 0th index of the 26 index "sCheck".
	//Since we are incrementing this and eventually going to be decrementing the same letter if this is in fact an Anagram, it should mean the entire array is going to be zeroed out. Otherwise we will return a false.
	for i := 0; i < sl; i++ {
		sCheck[s[i]-'a']++
		sCheck[t[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if sCheck[i] != 0 {
			return false
		}
	}

	return true

}
