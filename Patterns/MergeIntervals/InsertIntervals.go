package mergeintervals

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