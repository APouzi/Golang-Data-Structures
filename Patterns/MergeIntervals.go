package patterns

import (
	"fmt"

	sorting "github.com/APouzi/go-algos/Sorting"
)

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
	var sortedIntervals [][]int = MergeSortIntervalByStart(intervals, 0, len(intervals)-1)
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
	var sortedIntervals [][]int = MergeSortIntervalByStart(intervals, 0, len(intervals)-1) //{{1,2},{1,3},{2,3},{3,4},{4,6},{5,6}}
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


// You are given a 2D integer array intervals, where intervals[i] = [lefti, righti] describes the ith interval starting at lefti and ending at righti (inclusive). The size of an interval is defined as the number of integers it contains, or more formally righti - lefti + 1. 
//You are also given an integer array queries. The answer to the jth query is the size of the smallest interval i such that lefti <= queries[j] <= righti. If no such interval exists, the answer is -1.

// Return an array containing the answers to the queries.

// Example 1:
// Input: intervals = [[1,4],[2,4],[3,6],[4,4]], queries = [2,3,4,5]
// Output: [3,3,1,4]
// Explanation: The queries are processed as follows:
// - Query = 2: The interval [2,4] is the smallest interval containing 2. The answer is 4 - 2 + 1 = 3.
// - Query = 3: The interval [2,4] is the smallest interval containing 3. The answer is 4 - 2 + 1 = 3.
// - Query = 4: The interval [4,4] is the smallest interval containing 4. The answer is 4 - 4 + 1 = 1.
// - Query = 5: The interval [3,6] is the smallest interval containing 5. The answer is 6 - 3 + 1 = 4.

// Example 2:
// Input: intervals = [[2,3],[2,5],[1,8],[20,25]], queries = [2,19,5,22]
// Output: [2,-1,4,6]
// Explanation: The queries are processed as follows:
// - Query = 2: The interval [2,3] is the smallest interval containing 2. The answer is 3 - 2 + 1 = 2.
// - Query = 19: None of the intervals contain 19. The answer is -1.
// - Query = 5: The interval [2,5] is the smallest interval containing 5. The answer is 5 - 2 + 1 = 4.
// - Query = 22: The interval [20,25] is the smallest interval containing 22. The answer is 25 - 20 + 1 = 6.

// This is extremely similar to 213. The Skyline Problem, The only difference here is that:

// height = width for each building (compared to random height)
// min of all overlapped building (compared to max / skyline of overlapped building)
// So remember this type of problem using heap:)

func MinInterval(intervals [][]int, queries []int) []int {
	//Like most problems, we need to sort our intervals doing a stable sort, usually merge sort. This allows us to easily see the intervals in a contiguous fashion, seeing any overlapping and allow us to easily compare who where the smart is smaller or bigger.
    sortedIntervals := MergeSortIntervalByStart(intervals,0,len(intervals)-1)

	//QueryCopy is going to serve two purposes, one is that it will allow to keep the order of the original query pair, Which is why we are going to be appened to result, and also we need to sort this because of the fact that we need to easily keep track of the start of the intervals. This is how we know that we are going to be in the appropriate interval without having to brute force the solution. Example: say we have 4,7,2 and we have [(1,3)(1,2),(2,3)]. We can see how would have to be brute forcing our way into the solution. 
	queryCopy := make([]int, len(queries))
	copy(queryCopy,queries)
	sorting.QuickSort(queryCopy,0,len(queries)-1)//Now we need to sort the queries, this is what makes the sorting easier 
	//This is a minHeap that is sorting based on the size of the interval and also keeping the right most of the interval (end of interval)
	h := &HeapPairMin{}
	result :=map[int]int{}
	//The i represents the interation through the intervals
	i := 0
	//Here we are going to be looping through the sorted query. 
	for q := 0; q < len(queries);q++{
		//For every query, we need to insert every single interval that houses the query (query is bigger than or equal to start of interval). This is why we sorted the intervals. For each interval that this is true, we need to insert the range of said query and also the right most of the interval (more on this on 2nd loop.). Every single interval that is added to the heap are all possible canidates of the query and only the heap will be able to tell us which one is the best based on the range. 
		for i < len(sortedIntervals) && sortedIntervals[i][0] <= queryCopy[q]{
			h.Insert([]int{sortedIntervals[i][1] - sortedIntervals[i][0] + 1,sortedIntervals[i][1]})
			i++
		}
		//Once we added every single interval that houses the query, we need to start popping them based on the fact that end of each interval is smaller than the query. The reason for this is that this is asking is that we need to remove all those where the query is past the end, because it is no longer part of the intervals inside the range. The reason for this, is that if this loop is ever activated, it means that the we have moved onto a new query and the intervals that were added on a previous iteration no longer apply to this one. 
		for len(h.arr) >0 && h.arr[0][1] < queryCopy[q]{
			h.Pop()
		}
		//Check if the heap is empty, if not we need to add the range that is the range that is the heap.
		if len(h.arr) > 0{
			result[queryCopy[q]] =  h.arr[0][0]
		}else{
			result[queryCopy[q]] = -1
		}
	}
	var newList []int = make([]int,len(queries))
	for i := 0; i<len(queries);i++{
		newList[i] = result[queries[i]]
	}
	return newList
}


//Given an array of meeting time intervals consisting of start and end times[[s1,e1],[s2,e2],...](si< ei), determine if a person could attend all meetings.
//Input:[[0,30],[5,10],[15,20]]
//Output: false
func MeetingRooms(meetings [][]int)bool{
	MergeSortIntervalByStart(meetings, 0, len(meetings)-1)
	lastMeeting := meetings[0][1]
	for i := 1; i<len(meetings);i++{
		if lastMeeting >= meetings[i][0]{
			return false
		}
		lastMeeting = meetings[i][0]
	}
	return true
}


type Event struct{
	Time int
	Start bool
}

//Given an array of meeting time intervals consisting of start and end times[[s1,e1],[s2,e2],...](si< ei), find the minimum number of conference rooms required.
func MeetingRoomsII(intervals [][]int) int{
	var events []Event = make([]Event, len(intervals)*2)
	for i :=0;i<len(intervals);i++{
//Here we are turning the intervals into an event sequence list and now we are turning the list into 1D list. This is going to be allowing us to give us the imaginary sweep line to activate events such as start, end, overlapping and activation. We also input whether this specific Interval time is a start or not. Other algorithms where we want to know if there is active segments.
		events[2*i] = Event{Time:intervals[i][0], Start: true}
		events[2*i+1] = Event{Time:intervals[i][1], Start: false}
	}
	//List: {0, 30},{5, 10},{15, 20}
	//Events: {0 true},{5 true},{10 false},{15 true},{20 false},{30 false}
	MergeSortEvent(events,0,len(events)-1)
	var count,ans int = 0, 0 
	for i :=0; i < len(events);i++{
		if events[i].Start{ //Because of the fact that we sorted the Events, it would mean that everything is sorted and whats good is that because of this, we can keep track of the count overlapping.
			count++
			ans = max(count, ans)
		}else{
			count--
		}
	}
	return ans
}

func MeetingRoomsIIHeap(intervals [][]int)int{
	var maxAns int = 0
	active := InitializeHeapPairMin()
	MergeSortIntervalByStart(intervals,0,len(intervals)-1)
	for i :=0;i<len(intervals);i++{
		fmt.Println(active.arr, intervals[i])
		for len(active.arr) > 0 && active.arr[0][1] <= intervals[i][0]{
			active.Pop()
		}
		maxAns = max(maxAns, len(active.arr))
	}
	return maxAns
}


// You are given a 2D integer array logs where each logs[i] = [birthi, deathi] indicates the birth and death years of the ith person.

// The population of some year x is the number of people alive during that year. The ith person is counted in year x's population if x is in the inclusive range [birthi, deathi - 1]. Note that the person is not counted in the year that they die.

// Return the earliest year with the maximum population.

//Example 1:
// Input: logs = [[1993,1999],[2000,2010]]
// Output: 1993
// Explanation: The maximum population is 1, and 1993 is the earliest year with this population.

// Example 2:
// Input: logs = [[1950,1961],[1960,1971],[1970,1981]]
// Output: 1960
// Explanation: 
// The maximum population is 2, and it had happened in years 1960 and 1970.
// The earlier year between them is 1960.

func MaximumPopulation(logs [][]int) int {

	var events []Event = make([]Event, len(logs)*2)
	for i := 0; i < len(logs);i++{//Lets start creating the events that will be "activating" segments to allow us to count the maximum population. Since we will be sorting these events, we can keep track of active by counting the starts and decrementing the starts. 
		events[2*i] = Event{Time:logs[i][0], Start: true}
		events[2*i+1] = Event{Time:logs[i][1], Start: false}
	}
	MergeSortEvent(events,0,len(events)-1)
	var pop, MaxPop int = 0,0 //This is the data structure we going to be using to keep track of "active". We do this by the fact that we are going to be keeping track of the starts by incrementing up and ends by counting down. When we reach a new max, we are going to activate earliestYear by setting it to that year.
	var earliestYear int = 0
	for i := 0; i <len(events);i++{
		if events[i].Start{
			pop++ 
			if pop > MaxPop{//We kept counting the population going up, now that we have a new max, it means that we assign that year, as it would be the earliest because we already sorted it as the new earliestYear. MaxPop also needs to be updated for bookkeeping, incase a bigger year has been reached.  
				MaxPop = pop
				earliestYear = events[i].Time
			}
		}else{//Decrement if the end has been reached.
			pop--
		} 
	}
	return earliestYear

}

//A city's skyline is the outer contour of the silhouette formed by all the buildings in that city when viewed from a distance. Given the locations and heights of all the buildings, return the skyline formed by these buildings collectively.

// The geometric information of each building is given in the array buildings where: 
//buildings[i] = [lefti, righti, heighti]:
// lefti is the x coordinate of the left edge of the ith building.
// righti is the x coordinate of the right edge of the ith building.
// heighti is the height of the ith building.

// You may assume all buildings are perfect rectangles grounded on an absolutely flat surface at height 0.
// The skyline should be represented as a list of "key points" sorted by their x-coordinate in the form [[x1,y1],[x2,y2],...]. Each key point is the left endpoint of some horizontal segment in the skyline except the last point in the list, which always has a y-coordinate 0 and is used to mark the skyline's termination where the rightmost building ends. Any ground between the leftmost and rightmost buildings should be part of the skyline's contour.

// Note: There must be no consecutive horizontal lines of equal height in the output skyline. For instance, [...,[2 3],[4 5],[7 5],[11 5],[12 7],...] is not acceptable; the three lines of height 5 should be merged into one in the final output as such: [...,[2 3],[4 5],[12 7],...]

//Example 1
// buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
// Output: [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]

// Example 2:
// buildings = [[0,2,3],[2,5,3]]
// Output: [[0,3],[5,0]]

// func GetSkyline(buildings [][]int) [][]int {
// 	var skyline [][]int
// 	h := &HeapPairMax{}

// 	// Collect all x's (start and end of buildings)
// 	var xPlane []int = make([]int, len(buildings)*2)
// 	for i, b := range buildings {
// 		xPlane[i*2] = b[0]
// 		xPlane[i*2+1] = b[1]
// 	}

// 	//Sort all the newly made x points, we need to do this because we assume that the buildings are already sorted.
// 	sort.Ints(xPlane)
// 	var uniqueX []int = []int{xPlane[0]}
// 	//Now make sure that all x's are unique. We do this by the fact that each iteration will be related to the next in the sense that they are on a 1D plane because of the two x's. These need to be unique because conceputually, this problem will never have the skyline list the same X. When you think about a rectangle, you only need three points to actually create a rectangle ("think of the mouse drag on desktop")
// 	for i := 1; i < len(xPlane); i++ {
// 		if xPlane[i] != xPlane[i-1] {
// 			uniqueX = append(uniqueX, xPlane[i])
// 		}
// 	}

// 	//For every iteration of the uniqueX, we want to add all the segments that are needed into the heap, aka (active segments). We also use uniqueX the same way we use the queries is uniqueX, as way to range over possible "part of the answers" that will be lead to the final answer, an auxilaury look up. The only difference is that we need to actually build it out, then we have uniqueX.
// 	b := 0
// 	prevMaxHeight := 0
// 	fmt.Println("!!!!!!!!!!!",uniqueX)
// 	for _, x := range uniqueX {
// 		// Add buildings to the heap if their start is at or before x
// 		for b < len(buildings) && buildings[b][0] <= x {
// 			//Insert the building height and then insert the building right side (right x), we are going to be sorting this heap by the height. We need to insert the right in order for us to disqualify heights as we move onto the different X's.
// 			h.Insert([]int{buildings[b][2], buildings[b][1]})
// 			b++
// 		}
// 		//Remove buildings from the heap if their end is before x or equal to x. There can only be one x on the plane, left or right (mostly left)
// 		for len(h.arr) > 0 && h.arr[0][1] <= x {
// 			h.Pop()
// 		}
// 		//Get the current max height, notice that it starts at zero and if the heap has anything, the top of that heap will be the new current height. If the heap is empty, it means we have reached the floor and the currentHeight will stay 0.
// 		currentMaxHeight := 0
// 		if len(h.arr)  > 0 {
// 			currentMaxHeight = h.arr[0][0]
// 		}

// 		//Update the skyline if the max height changed, append it to the skyline and then assign the prevMaxHeight to currentMaxHeight
// 		if currentMaxHeight != prevMaxHeight {
// 			skyline = append(skyline, []int{x, currentMaxHeight})
// 			prevMaxHeight = currentMaxHeight
// 		}
// 	}

// 	return skyline
// }



func MergeIntervalsAttempt(intervals [][]int) [][]int {
	sortedIntervals := MergeSortIntervalByStart(intervals, 0, len(intervals)-1)
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