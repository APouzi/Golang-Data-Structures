package stack

// Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.

// Example 1:
// Input: temperatures = [73,74,75,71,69,72,76,73]
// Output: [1,1,4,2,1,1,0,0]

// Example 2:
// Input: temperatures = [30,40,50,60]
// Output: [1,1,1,0]

// Example 3:
// Input: temperatures = [30,60,90]
// Output: [1,1,0]
type Temperature struct{
	temp int
	index int
}
func dailyTemperatures(temperatures []int) []int {
    stack := []Temperature{}
    ans := make([]int,len(temperatures))
    for i := 0; i < len(temperatures);i++{
        for len(stack) > 0 && temperatures[i] > stack[len(stack)-1].temp{ //Monotonically decreasing means that bottom of the stack is the biggest and top is the smallest in the current iteration.
			ans[stack[len(stack)-1].index] = i - stack[len(stack)-1].index //Because we are storing the index, this is how we know that length of time that has passed as the iteration go on. 
            //Due to the fact that stack holds the iteration, this is how we keep track of when to insert the said answer to the correct position. 
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, Temperature{temperatures[i],i})
    }
    return ans
}


//This right here is what I would say is what's important. It's the raw version, where all we really care about the stack is the iteration, this allows us to keep track of range and this range being the "days".
func dailyTemperaturesIndex(temperatures []int) []int {
    stack := []int{}
    ans := make([]int,len(temperatures))
    for i := 0; i < len(temperatures);i++{
        for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]]{ 
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1] 
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    return ans
}