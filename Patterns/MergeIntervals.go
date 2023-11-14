package patterns

import "fmt"

//Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

// Example 1:
// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

// Example 2:
// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.

//Note, DO NOT use quicksort for intervals, this is because quicksort intervals is unstable.
func MergeIntervals(intervals [][]int) [][]int {
	var sortedIntervals [][]int = MergeSortInterval(intervals, 0, len(intervals)-1)
	fmt.Println(sortedIntervals)
	var merged [][]int = make([][]int,0)
	//No matter what, we need to add the 1st interval
	merged = append(merged, sortedIntervals[0])
	for i := 1; i<len(sortedIntervals);i++{
		//In this instance, we are asking if the pair's 2nd element is smaller than the 1st element in the next pair, then we add it. This just implies we have a disjointed set.
		if merged[len(merged)-1][1] < sortedIntervals[i][0]{
			merged =append(merged, sortedIntervals[i])
		}else{
			//Here we found that the two pairs are actually joined because the 2nd element in the 1st pair is in fact bigger than the 1st element in the next pair, but need to figure out if we should merge them by comparing both pair's 2nd elements and seeing which one is bigger. Remember that [1,4],[4,5] are overlapping
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1],sortedIntervals[i][1])
		}
	}
	return merged
}

//You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.
//Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).
//Return intervals after the insertion.
 
//Example 1:
//Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
//Output: [[1,5],[6,9]]

//Example 2:
//Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//Output: [[1,2],[3,10],[12,16]]
//Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].


func InsertIntervals(intervals [][]int, newInterval []int)[][]int{
	var listAns  [][]int = make([][]int, 0)

	for i := 0; i < len(intervals); i++{
		//Anything to the left of new interval?:If the 2nd element in the interval is smaller than the 1st element in the newInterval to be inserted, then that means this interval comes before the newInterval (the one to be inserted).
		if intervals[i][1] < newInterval[0] {
			listAns = append(listAns, intervals[i])
		//Anything to the right to the new interval?: Here we are saying that if intervals 2nd element is bigger than the newIntervals 1st element, that means all subsequent newPairs are done and there is no way for us to do anything else for the newInterval. So we add the newInterval 1st and then return the rest of the intervals.
		} else if intervals[i][0] > newInterval[1] {
			listAns = append(listAns, newInterval)
			return append(listAns, intervals[i:]...)
		//The actual insertion and modifying of newInterval:Lastly, because our newInterval to be inserted does not go after the current iteration interval and does not go before, this means that the only possible thing to do next is to start modifying the current interval to grow it insize and to do this until we are done with it. Which means we either reach the 1st if statement or we loop through and add it at the end of the list. We need to modify the new interval in place by asking which interval is smaller for the 1st element and which interval is bigger on the 2nd element
		} else {
			newInterval[0], newInterval[1] = min(newInterval[0], intervals[i][0]), max(newInterval[1], intervals[i][1])
		}
	}
	listAns = append(listAns, newInterval)
	return listAns
}


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
	var sortedIntervals [][]int = MergeSortInterval(intervals, 0, len(intervals)-1)
	fmt.Println(sortedIntervals)
	var lastInterval []int = sortedIntervals[0]
	for i := 1; i < len(sortedIntervals); i++{
		if lastInterval[1] <= sortedIntervals[i][0]{
			lastInterval[1] = sortedIntervals[i][1]
		}else{
			ans++
			lastInterval[1] = min(lastInterval[1], sortedIntervals[i][1])
		}
	}
	return ans
}
















func MergeIntervalsAttempt(intervals [][]int) [][]int {
	sortedIntervals := MergeSortInterval(intervals, 0, len(intervals)-1)
	fmt.Println(sortedIntervals)
	listAns := [][]int{}
	var newPair []int
	for i := 0; i<len(sortedIntervals)-1;i++{
		//(1,3),(2,6)
		if sortedIntervals[i][1] >= sortedIntervals[i+1][0]{
			newPair = []int{sortedIntervals[i][0],sortedIntervals[i+1][1]}
		}
		if len(listAns) > 0 && newPair[0] == listAns[len(listAns)-1][0] && listAns[len(listAns)-1][1] < newPair[1]{
			listAns = listAns[:len(listAns)-1]
			listAns = append(listAns, newPair)
		} else if len(listAns) == 0 || newPair[0] != listAns[len(listAns)-1][0]{
			listAns = append(listAns, newPair)
		}
	}
	return listAns
}