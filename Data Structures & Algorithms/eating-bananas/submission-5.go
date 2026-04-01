func minEatingSpeed(piles []int, h int) int {
	l,r := 1,0
	result := 0

	for _, p := range piles {
		r = max(r,p)
	}

	for l <= r {
		mid := l + (r-l)/2

		if canEatAllBanana(piles,mid,h) {
			result = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return result

}

func canEatAllBanana(piles []int, speed, h int) bool {
	hs := 0
	for _, p := range piles {
		hs += (p+speed-1)/speed
		
		if hs > h {
			return false
		}
	}

	return hs <= h
}
