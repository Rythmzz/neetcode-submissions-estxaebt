func minEatingSpeed(piles []int, h int) int {
	l,r := 1,0
	result := 0

	for _,p := range piles {
		r = max(r,p)
	}

	for l <= r {
		mid := l + (r-l)/2

		if eatAllBanana(piles,mid,h) {
			result = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return result
}

func eatAllBanana(piles []int, speed,h int) bool {
	hSpeed := 0

	for _,p := range piles {
		hSpeed += (p + speed - 1)/speed
		if hSpeed > h {
			return false
		}
	}

	return hSpeed <= h

}
