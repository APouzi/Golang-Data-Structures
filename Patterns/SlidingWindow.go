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
		// fmt.Println(dynamicBank, bankStr2)
		
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


//Given an array of integers nums and an integer k, return the number of contiguous subarrays where the product of all the elements in the subarray is strictly less than k.

// Example 1:
// Input: nums = [10,5,2,6], k = 100
// Output: 8
// Explanation: The 8 subarrays that have product less than 100 are:
// [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
// Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.

// Example 2:
// Input: nums = [1,2,3], k = 0
// Output: 0

func NumSubarrayProductLessThanK(nums []int, k int) int {
    var left int = 0 //This is the left pointer that will be used to shrink the window
    var product int = 1
    var ans int = 0
    for right := 0; right < len(nums); right++{ //The right pointer is used to grow the window
        product *= nums[right]//On start, we need to grow the window. given ex1, it grows to [10,5,2] which would activate the loop below. Now this is also the part where we start moving this window to the right. Once we shrink the window to the appropriate condition, we need to start growing it so it goes from [10,5,2] to [5,2] then [5,2,6]
        for left <= right && product >= k{//What happens here is that we need to shrink the window until either we shrink it to a single element or product becomes smaller than K
            product /= nums[left] //This is our window shrinking on removing the left side of the window
            left++
        }
        ans += (right - left) +1 //This is the confusing part. The possible windows are [10], [10,5], [5,2], [5,2,6], how are we getting the answers? Well first, we must internalize that we are getting the possible windows that are less than K or whatever our window goal is. 
		//Later, we are realizing that for every window we add, the new element that we add is it's own subarray. 
		//So an example "[10]" is one subarr, so +1, 
		//[10,5] is anther subarr (+1) plus the fact that [5] by itself is one too (+1), so we must add 2 this. 
		//[5,2] is another subarr(+1) plus another subarr which is [2] (+1) so we add 2. 
		//[5,2,6], which is a subarr so (+1), [2,6] so subarr (+1), and lastly [6] so (+1), this is why we add 3 on this one
    }
    return ans
}