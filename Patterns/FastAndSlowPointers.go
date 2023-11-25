package patterns

import "fmt"

// Write an algorithm to determine if a number n is happy.
// A happy number is a number defined by the following process:
// Starting with any positive integer, replace the number by the sum of the squares of its digits.Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.
func HappyNumber(n int) bool {
	SquareSumDigit := func(num int) int {
		currsum := 0
		for num > 0 {
			digit := num % 10
			currsum += digit * digit
			num /= 10
		}
		return currsum

	}

	var fast, slow int = n, n

	for {
		slow = SquareSumDigit(slow)
		fast = SquareSumDigit(SquareSumDigit(fast))

		if fast == slow {
			break
		}
	}

	return slow == 1
}


// This is just a placeholder, we are not actually running this because it's currently inside of the linkedlist
type LinkedNode struct{
	Next *LinkedNode
	Val string
}
func CycleDetect(node *LinkedNode) bool {
	fast := node
	slow := node

	for fast != nil || slow != nil { //Here we are stopping if the linked list is a not a cycle
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow { //If it is a cycle we are gauranteed to stop because eventually this will become true
			fmt.Println("connected at: ", fast.Val)
			return true
		}
	}
	return false
}

//Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

// There is only one repeated number in nums, return this repeated number.
// You must solve the problem without modifying the array nums and uses only constant extra space.

// Example 1:
// Input: nums = [1,3,4,2,2]
// Output: 2

// Example 2:
// Input: nums = [3,1,3,4,2]
// Output: 3

//Example 3:
// Input: nums = [2,5,9,6,9,3,8,9,7,1]
// Output: 9
//get the true index:
// 9 5, 1 6, 5 7, 3 1, 6 3, 8 8, 
// get the Number:
// 7 2, 9 9, 
func FindDuplicate(nums []int) int {
    var slow, fast int = nums[0], nums[nums[0]]
// These array's are only representing the next index that it goes to, because these are cyclical arrays that basically point to the next index. So calling itself will allow us to infinitely call the array without ever getting out of bounds

    for slow != fast{ //On the first loop, we are simply following the array's path to get the "cycle's entry point".
        slow = nums[slow]
        fast = nums[nums[fast]]
		//slow[1]:     3,    2,    |4|
		//fast[3]: [2]>4,[2]>4,[2]>|4|
    }
    fast = 0
    for slow != fast{//After we found the "entry point into the cycle" index (first example being the 4th index), now need to find this array's index
        slow = nums[slow]
        fast = nums[fast]
		//slow(4):2,4,|2|
		//fast(0):1,3,|2|
		
    }
    return fast
}