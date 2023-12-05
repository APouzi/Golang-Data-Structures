package patterns

import (
	"fmt"

	fastandslow "github.com/APouzi/go-algos/Patterns/FastAndSlow"
	slidingwindow "github.com/APouzi/go-algos/Patterns/SlidingWindow"
)

func RunPattern() {
	// ---Permutation String---
	returnRes := slidingwindow.PermutationOfString("eidboaoo", "eidboaoo")
	fmt.Println("PermutationOfString",returnRes)
	//"eidboaoo", str2 "eidboaoo"
	returnpracRes := PermutationOfStringPrac("eidboaoo", "ba")
	fmt.Println("PermuStrPrac result",returnpracRes)

	// --- High Sell low
	returnRes2 := slidingwindow.BuyHighSellLow([]int{6, -5, 4, 6, 1, 2, 5, 6, 5, 4, 6, 1, 2, 5, 6, 5, 4, 6, 1, 2, 11})
	fmt.Println(returnRes2)

	//----SubArray Product less than K---
	maxProdArr := []int{10,5,2,6}
	fmt.Println("Number SubArray Product less than K",slidingwindow.NumSubarrayProductLessThanK(maxProdArr,100))


	// ---Happy Number---
	HappyInput := 19
	HappyNum := fastandslow.HappyNumber(HappyInput)
	fmt.Printf("Is the number \"%d\", a happy num? Answer: %t\n", HappyInput,HappyNum)
	HappyNumPrac := HappyNumberPrac(HappyInput)
	fmt.Printf("Practice! Is the number \"%d\", a happy num? Answer: %t\n", HappyInput,HappyNumPrac)


	//---Merge Intervals---
	intervals := [][]int{{2,5},{1,3},{1,4},{9,15},{6,10}, {7,13}}
	fmt.Println("Merge Intervals:",MergeIntervals(intervals))

	intervalsPrac := [][]int{{2,5},{1,3},{1,4},{9,15},{6,10}, {7,13}}
	fmt.Println("Merge Intervals Practice:",MergeIntervalsPrac(intervalsPrac))

	//---Insert Intervals---
	intervalsIns := [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
	fmt.Println("Insert Intervals:",InsertIntervals(intervalsIns,[]int{4,8}))

	intervalsInsPrac := [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
	fmt.Println("Insert Intervals Prac:",InsertIntervalsPrac(intervalsInsPrac,[]int{4,8}))

	//---Non OverLapping Intervals---
	intervalsRem := [][]int{{1,2},{2,3},{3,4},{1,3}}
	fmt.Println("Non OverLapping Intervals:", nonOverLappingIntervals(intervalsRem))

	intervalsRem2 := [][]int{{1,2},{2,3},{3,4},{1,3}}
	fmt.Println("Non OverLapping Intervals Prac:", nonOverLappingIntervalsPrac(intervalsRem2)) 

	//---Minimum Interval to Include Each Query---
	// intervalsMin := [][]int{{1,4},{2,4},{3,6},{4,4}} 
	// queries := []int{2,3,4,5}
	intervalsMin := [][]int{{2,3},{2,5},{1,8},{20,25}}
	queries := []int{2,19,5,22}
	// intervalsMin := [][]int{{4,5},{5,8},{1,9},{8,10},{1,6}}
	// queries := []int{7,9,3,9,3}
	fmt.Println("Minimum Interval to Include Each Query:", MinInterval(intervalsMin,queries))
	fmt.Println("Minimum Interval to Include Each Query Practice:", MinIntervalPrac(intervalsMin,queries))

	//---Meeting Rooms and Meeting Rooms II----
	meetings := [][]int{{0,30},{5,10},{15,20}}
	// meetings := [][]int{{7,10},{2,4}}
	fmt.Println("Meeting Rooms:", MeetingRooms(meetings))

	// meetingsII := [][]int{{0, 30},{5, 10},{15, 20}}
	meetingsII := [][]int{{7,10},{2,4}}
	fmt.Println("Meeting Rooms II:",MeetingRoomsII(meetingsII))
	fmt.Println("Meeting Roomms II",MeetingRoomsIIHeap(meetingsII))

	//---Maximum Year---
	// Years := [][]int{{1950,1961},{1960,1971},{1970,1981}}
	Years := [][]int{{2008,2026},{2004,2008},{2034,2035},{1999,2050},{2049,2050},{2011,2035},{1966,2033},{2044,2049}}
	fmt.Println("Maximum Years:", MaximumPopulation(Years))

	//---SkyLine Problem---
	buildings := [][]int{{2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}}
	fmt.Println("SkyLine", GetSkyline(buildings))
	fmt.Println("Skyline Practice:",GetSkylinePrac(buildings))

	//---Fruit Basket---
	fmt.Println("Fruit Basket Practice", TotalFruitPrac([]int{1,2,3,2,2}))

	//---Character Replacement---
	fmt.Println("Longest Repeating Characters With Repeating Characters", CharacterReplacementPrac("AABABBA",2))
	

}