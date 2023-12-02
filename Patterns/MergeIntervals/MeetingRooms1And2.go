package mergeintervals

import "fmt"

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