package binarysearch

import "math"

// Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

// Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

// Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

// Return the minimum integer k such that she can eat all the bananas within h hours.

// Example 1:
// Input: piles = [3,6,7,11], h = 8
// Output: 4

// Example 2:
// Input: piles = [30,11,23,4,20], h = 5
// Output: 30

// Example 3:
// Input: piles = [30,11,23,4,20], h = 6
// Output: 23

func MinEatingSpeed(piles []int, h int) int {
	//Step one is to create an options list that we are going to do a binary search on. This options list will be the minimum amount koko can eat from the pile, 1. To the max amount. This list is how we can perform the binary search to find out the minimum amount of bananas we can eat per hour, without surpassing the hour.
    var left, right = 1, maxPilesSize(piles)
    var result int = right
    var middle int
    var hours float64 = 0
    h64 := float64(h)
    for left <= right{
        middle = (right + left)/2 //The middle, otherwise known as "k", is the canidate we have to see if this we add up the time, it's less than the hours (h) given. 
        hours = 0
        for _, v := range piles{ //Calculate if the canidate eating rate will fit into the given hours before gaurds come back.
            hours += math.Ceil(float64(v)/float64(middle))
        }

        if hours <= h64 { //If the hours are less than the hours we have before gaurds come back, this means we can lessen the possible canidates by half aka larger the middle. We also want to continually see if that's the minimum possible to eat. Because koko wants the least as possible.
            result = min(result, middle)
            right = middle -1
        }else if hours >= h64 { // If the hours given is too big, it means we can lessen by half, but starting from the left. This is because of the fact that due to the fact that the eat rate, as in the middle is too small and we need to increase the eat rate!
            left = middle + 1
        }
    }

    return result
}


func MinEatingSpeedV2(piles []int, h int) int {
	
	var left, right = 1, maxPilesSize(piles)
	var result int = right
	var middle int
	var hours int = 0
	for left <= right {
		middle = (right + left) / 2
		hours = 0
		for _, v := range piles {
			hours += ((v + middle) - 1) / middle
		}

		if hours <= h {
			result = min(result, middle)
			right = middle - 1
		} else if hours >= h {
			left = middle + 1
		}
	}

	return result
}

func maxPilesSize(piles []int) int {
	var ans int = 0
	for _, v := range piles {
		ans = max(ans, v)
	}
	return ans
}