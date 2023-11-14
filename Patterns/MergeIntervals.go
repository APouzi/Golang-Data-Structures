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
	//No matter what, we need to add the first interval
	merged = append(merged, sortedIntervals[0])
	for i := 1; i<len(sortedIntervals);i++{
		//In this instance, we are asking if the pair's second element is smaller than the first element in the next pair, then we add it. This just implies we have a disjointed set.
		if merged[len(merged)-1][1] < sortedIntervals[i][0]{
			merged =append(merged, sortedIntervals[i])
		}else{
			//Here we found that the two pairs are actually joined because the 2nd element in the first pair is in fact bigger than the first element in the next pair, but need to figure out if we should merge them by comparing both pair's 2nd elements and seeing which one is bigger. Remember that [1,4],[4,5] are overlapping
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1],sortedIntervals[i][1])
		}
	}
	return merged
}

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
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