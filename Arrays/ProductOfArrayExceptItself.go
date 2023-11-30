package arrays

//Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i]. The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

// Example 2:
// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]


func ProductExceptSelf(nums []int) []int {
	var prefix, postfix, n int = 1, 1, len(nums) - 1 // What we are going to be doing is doing two passes over the result. One for the prefix, as we move forward and one for the postfix as we move back. We need to assign everyone of these of 1 because of the first and last position being zeroed out.
	var res []int = make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = prefix // [1,1,2,6]
						//prefix * nums[i]
		prefix *= nums[i] //1*1, 1*2, 2*3, 3*4 (last one doesn't get added)
	}
	// fmt.Println("Prefix", res)
	for i := n; i >= 0; i-- {
		res[i] *= postfix //6*1,4*2,12*1, [24,12,8,6](remember, we are adding to this list in backwards order)
		postfix *= nums[i]//1*4, 4*3, 12*2, 24*1
	}
	return res
}

//nums =   [1,2,3,4]
//prefix = [1,1,2,6]