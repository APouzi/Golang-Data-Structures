package patterns

import "fmt"

//Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.
// Example 1:
// Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
// Output: 1
// Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.

// Example 2:
// Input: intervals = [[1,2],[1,2],[1,2]]
// Output: 2
// Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.

// Example 3:
// Input: intervals = [[1,2],[2,3]]
// Output: 0
// Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
//NOTE as example 3 shows, if two edges are the same, they are NOT overlapping.
func nonOverLappingIntervals(intervals [][]int) int{
	var ans int = 0
	var sortedIntervals [][]int = MergeSortIntervalByStart(intervals, 0, len(intervals)-1) //{{1,2},{1,3},{2,3},{3,4},{4,6},{5,6}}
	fmt.Println(sortedIntervals)
	var lastInterval []int = sortedIntervals[0] //NOTE: this doesn't need to be interval and we can just keep track of the last good end, as we shrink this.
	for i := 1; i < len(sortedIntervals); i++{
		if lastInterval[1] <= sortedIntervals[i][0]{
			lastInterval[1] = sortedIntervals[i][1] //Here, we are saying that we are going to be growing the interval to the next intervals end IF the last intervals end is smaller or equal to the current iteration. This is because that means it's not overlapping.
		}else{
	//If the last interval's end is bigger/overlapping the start of the next interval, then we want to increment the number of removal. The last bit of this is that we want to shorten the intercal to the minimum end of the two intervals, in "transforming it" into the previous interval. What this does is when we move onto the next interval, we are not going to trigger the response of needing to delete the overlapping because it's no longer needed. 
			ans++
			lastInterval[1] = min(lastInterval[1], sortedIntervals[i][1])
	//Given this example: {{1,2},{1,3},{2,3},{3,4},{4,6},{5,6}}, what happens here is we start with (1,2) as last interval, then we increment result to 1 due to end of (1,2) overlapping (1,3). We then grow the last interval shrink "(1,3)" to "(1,2)". After this, we will grow (1,2) all the way to (1,6) because that is the last possible interval that doesn't overlap and we need this to not crawl up due to the fact that (4,6) to (5,6) are overlapping and we would need to delete one of this, being minimum.
		}
	}
	return ans
}