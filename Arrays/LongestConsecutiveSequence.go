package arrays

//Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
// You must write an algorithm that runs in O(n) time.

// Example 1:
// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

// Example 2:
// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9

func LongestConsecutive(nums []int) int {
    var consec map[int]bool = make(map[int]bool)

    for _, v:= range nums{
        consec[v] = true
    }
    var ans int =0
    for i := 0; i < len(nums);i++{
        if consec[nums[i]-1]{
            continue
        }
        var count int = 1
        var check int = nums[i] +1
        for consec[check]{
            count++
            check++
            
        }
        ans = max(count, ans)
    }
    return ans
}

func max(i, j int) int{
	if i > j{
		return i
	}
	return j
}

