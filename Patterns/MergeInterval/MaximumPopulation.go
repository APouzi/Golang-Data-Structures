package patterns

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