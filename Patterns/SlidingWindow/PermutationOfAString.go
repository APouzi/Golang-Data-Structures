package slidingwindow

import (
	"reflect"
)

// Given a string(str1), is str2 a permutation of str1?
// Example str1 = "abcdefg", str2 = "bac". "bac" is a permuatation of "abcdefg" because "abc" of "abcdefg" can be permutated to "bac"
// test cases:
// str1 = "eidbaooo", str2 = "ab"
// str1 = "eidboaoo", str2 = "ab"
// str1 = "eidboaoo", str2 "eidboaoo"

func PermutationOfString(str1, str2 string) bool{
	compareTo := map[byte]int{}
	var need, have int = 0, 0
	for i := 0; i <len(str2);i++{
		compareTo[str2[i]]++
		have++	
	}

	bank := map[byte]int{}
	left := 0
	for r := 0; r <len(str1);r++{
		bank[str1[r]]++
		need++

		
		for need == have && left <= r{
			if reflect.DeepEqual(bank, compareTo){
				return true
			}
			
			bank[str1[left]]--
			if bank[str1[left]] == 0{
				delete(bank, str1[left])
			}
			left++
			need--
		}
	}
	return false
}