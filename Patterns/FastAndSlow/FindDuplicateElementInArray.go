package fastandslow

func FindDuplicate(nums []int) int {
	var slow, fast int = nums[0], nums[nums[0]]
	// These array's are only representing the next index that it goes to, because these are cyclical arrays that basically point to the next index. So calling itself will allow us to infinitely call the array without ever getting out of bounds. Which is why we first must find the "entry point" of the cycle with the cycle detection method.

	for slow != fast { //On the first loop, we are simply following the array's path to get the "cycle's entry point".
		slow = nums[slow]
		fast = nums[nums[fast]]
		//slow[1]:     3,    2,    |4|
		//fast[3]: [2]>4,[2]>4,[2]>|4|
	}
	fast = 0           //The reason we are setting fast (or slow) to the 0 is because the 0 will be outside of the cycles bound. What happens is that which ever one we didn't set to 0, is in the cycle loop. What it also means is that because of the fact that we are going to be setting it to zero, mathmatically speaking, the steps to enter the cycle is going to be the same amount of steps
	for slow != fast { //After we found the "entry point into the cycle" index (first example being the 4th index), now need to find this array's index
		slow = nums[slow]
		fast = nums[fast]
		//slow(4):2,4,|2|
		//fast(0):1,3,|2|

	}
	return fast
}