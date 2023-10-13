package patterns

import (
	"fmt"
	"reflect"
)

// Given a string(str1), is str2 a permutation of str1?
// Example str1 = "abcdefg", str2 = "bac". "bac" is a permuatation of "abcdefg" because "abc" of "abcdefg" can be permutated to "bac"
// test cases:
// str1 = "eidbaooo", str2 = "ab"
// str1 = "eidboaoo", str2 = "ab"
// str1 = "eidboaoo", str2 "eidboaoo"

func PermutationOfString(str1, str2 string) bool{
	dynamicBank := make(map[string]int)
	bankStr2 := make(map[string]int)

	for _, v := range str2{
		bankStr2[string(v)]+=1
	}
	winSize:=0

	for i, v := range str1 {
		if _,ok := dynamicBank[string(v)];ok{
			dynamicBank[string(v)] += 1
			winSize+=1
		}else{
			dynamicBank[string(v)]=1
			winSize+=1
		}
		fmt.Println(dynamicBank, bankStr2)
		
		if len(dynamicBank) >= len(bankStr2){

			// This will need the sliding window portion, shouldn't delete if key has above 2
			if winSize > len(str2){
				fmt.Println("hit")
				deleteThis := i-len(str2)
				if dynamicBank[string(str1[deleteThis])] > 1{
					dynamicBank[string(str1[deleteThis])] -=1
					winSize-=1
					fmt.Println("inside",dynamicBank)
					
				}else if dynamicBank[string(str1[deleteThis])] == 1{
					delete(dynamicBank,string(str1[deleteThis]))
					winSize-=1
				}				
			}
			if reflect.DeepEqual(dynamicBank, bankStr2){
				return true
			}
		}
	}

	return false
}

func BuyHighSellLow(prices []int) int{
	lowest := prices[0] //Originally did -999999, but that would mean that we could possibly get a worse conidition
	maxSell := -9999999999

	for _, v := range prices {
		possibleMS := v - lowest
		if lowest > v {
			lowest = v
		}

		if possibleMS > maxSell{
			maxSell = possibleMS
		}
	}
	return maxSell
}