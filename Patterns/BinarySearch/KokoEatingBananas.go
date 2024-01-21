package binarysearch

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