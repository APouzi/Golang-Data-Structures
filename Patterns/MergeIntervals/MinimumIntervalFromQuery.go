package mergeintervals

import sorting "github.com/APouzi/go-algos/Sorting"

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