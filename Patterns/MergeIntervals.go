package patterns

import "fmt"

//Note, DO NOT use quicksort for intervals, this is because quicksort intervals is unstable.
func MergeIntervals(intervals [][]int) [][]int {
	sortedIntervals := MergeSortInterval(intervals, 0, len(intervals)-1)
	fmt.Println(sortedIntervals)
	merged := [][]int{}
	// var newPair []int
	for i := 0; i<len(sortedIntervals);i++{
		//(1,3),(2,6)
		if len(merged) == 0 || merged[len(merged)-1][1] < sortedIntervals[i][0]{
			// newPair = []int{sortedIntervals[i][0],sortedIntervals[i+1][1]}
			merged =append(merged, sortedIntervals[i])
		}else{
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1],sortedIntervals[i][1])
		}
		// if sortedIntervals[i][1] >= sortedIntervals[i+1][0]{
		// 	// newPair = []int{sortedIntervals[i][0],sortedIntervals[i+1][1]}
		// }
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