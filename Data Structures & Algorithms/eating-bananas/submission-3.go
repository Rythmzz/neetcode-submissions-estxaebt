func minEatingSpeed(piles []int, h int) int {
	minSpeed := 0
	l,r := 1, 0

	for _, p:= range piles {
		r = max(r,p)
	}

	for l <= r {
		mid := l + (r-l)/2

		if canEatAllBanana(piles,mid,h) {
			minSpeed = mid
			r = mid - 1
		} else {
			l = mid + 1
		}

	}

	return minSpeed

}

func canEatAllBanana(piles []int, speed, h int) bool {
	mh := 0
	for _,p := range piles {
		mh += (p + speed - 1 )/speed
		if mh > h {
			return false
		}
	}

	return mh <= h
}
