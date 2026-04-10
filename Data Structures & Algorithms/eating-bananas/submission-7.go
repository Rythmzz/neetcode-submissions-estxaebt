func minEatingSpeed(piles []int, h int) int {
	l, r:= 1,0
	maxSpeed := 0
	for _, p:= range piles {
		r = max(p,r)
	}

	for l <= r {
		mid := l + (r-l)/2

		if canEatAllBanana(piles,mid,h) {
			maxSpeed = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return maxSpeed
}

func canEatAllBanana(piles []int, speed int, h int) bool {
	hoursUse := 0
	for _, p := range piles {
		hoursUse += (p + speed - 1)/speed

		if hoursUse > h {
			return false
		}
	}

	return hoursUse <= h
}
