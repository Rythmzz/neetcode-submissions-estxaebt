func minEatingSpeed(piles []int, h int) int {
	l, r := 1,0
	result := 0

	for _,p:= range piles {
		r = max(r,p)
	}

	for l <= r {
		mid := l + (r-l)/2

		if canAllEating(piles,mid,h) {
			result = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return result

	
}

func canAllEating(piles []int, speed int, h int) bool{
	hours := 0
	for _, p:= range piles{
		hours += (p + speed - 1)/ speed
	}

	if hours <= h {
		return true
	}

	return false
}
