package mergeintervals

import "sort"

//You are given two lists of closed intervals, firstList and secondList, where firstList[i] = [starti, endi] and secondList[j] = [startj, endj]. Each list of intervals is pairwise disjoint and in sorted order.

// Return the intersection of these two interval lists.

// A closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.

// The intersection of two closed intervals is a set of real numbers that are either empty or represented as a closed interval. For example, the intersection of [1, 3] and [2, 4] is [2, 3].

// Example 1:
// Input: firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
// Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]

// Example 2:
// Input: firstList = [[1,3],[5,9]], secondList = []
// Output: []


func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var newList [][2]int = make([][2]int, (len(firstList)+len(secondList))*2)
	firstList = append(firstList, secondList...)
	for i, v := range firstList {
		newList[i*2] = [2]int{v[0], 1}
		newList[i*2+1] = [2]int{v[1], -1}
	}

	sort.Slice(newList, func(i, j int) bool {
		if newList[i][0] == newList[j][0] {
			return newList[i][1] > newList[j][1]
		}
		return newList[i][0] < newList[j][0]
	})

	var ans [][]int = [][]int{}
	var startPoint, currentDelta int = 0, 0
	for i := 0; i < len(newList); i++ {
		currentDelta += newList[i][1]

		if currentDelta == 2 {
			startPoint = newList[i][0]
		}

		if currentDelta == 1 && newList[i][1] == -1 {
			ans = append(ans, []int{startPoint, newList[i][0]})
			startPoint = 0
		}

	
	}

	return ans
}

	// if newList[i-1][1] >= newList[i][0]{
		//     //Adjust the newList interval itself to account for intervals that have been changed.
		//     ans = append(ans, []int{max(newList[i][0],newList[i-1][0]),min(newList[i][1],newList[i-1][1])})
		// }